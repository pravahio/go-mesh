package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	grpcweb "github.com/improbable-eng/grpc-web/go/grpcweb"
	logging "github.com/ipfs/go-log"
	"github.com/pravahio/go-auth-provider/store"
	config "github.com/pravahio/go-mesh/config"
	mesh "github.com/pravahio/go-mesh/mesh"
	rpc "github.com/pravahio/go-mesh/rpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	log = logging.Logger("rpc-server")
)

// RPCServer brings compatibility for raw (http/2) and web (http/1.1) RPC
type RPCServer interface {
	Serve(net.Listener) error
	Stop()
}

// LisServer is a container for RPC Listener and Server
type LisServer struct {
	Listener net.Listener
	Server   RPCServer
}

// Server is the RPC server and contains variables needed for RPC to work.
type Server struct {
	m          *mesh.Mesh
	lmap       map[string]*LisServer
	serverOpts []grpc.ServerOption
	val        *store.Validator
}

// NewServer creates a new RPC server.
func NewServer(m *mesh.Mesh, authCrt string) (*Server, error) {

	v, err := store.NewValidator(authCrt)
	if err != nil {
		return nil, err
	}

	s := Server{
		m:    m,
		lmap: make(map[string]*LisServer),
		val:  v,
	}

	err = s.registerRPC(m.Cfg.RPC)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Server) registerRPC(rpcConfig []config.RPCConfig) error {
	if len(rpcConfig) > 0 {
		err := s.buildServerOptions(rpcConfig[0])
		if err != nil {
			return err
		}
	}
	for _, r := range rpcConfig {
		log.Info("RPC End: ", r.Endpoint)
		lis, err := net.Listen("tcp", r.Endpoint)
		if err != nil {
			s.CleanUp()
			return err
		}
		s.lmap[r.Mode] = &LisServer{
			Listener: lis,
		}
	}

	return nil
}

func (s *Server) buildServerOptions(config config.RPCConfig) error {
	if config.KeyPath == "" || config.CertPath == "" {
		s.serverOpts = []grpc.ServerOption{}
		log.Info("Both cert and key file paths are empty. Enabling insecure mode.")
		return nil // Create an insecure connection
	}
	cred, err := credentials.NewServerTLSFromFile(config.CertPath, config.KeyPath)
	if err != nil {
		return errors.New("Err in Credentials file: " + err.Error())
	}
	s.serverOpts = []grpc.ServerOption{
		grpc.UnaryInterceptor(s.handleUnaryCall),
		grpc.StreamInterceptor(s.handleStreamCall),
		grpc.Creds(cred),
	}
	return nil
}

func (s *Server) Start() {
	grpcServer := grpc.NewServer(s.serverOpts...)
	rpc.RegisterMeshServer(grpcServer, s)

	for m, ls := range s.lmap {
		dyServer := buildServer(grpcServer, m)
		if dyServer == nil {
			log.Warning("RPC Server is nil. Check the mode you passed.")
			continue
		}

		ls.Server = dyServer

		log.Info("Starting serveServer")
		serveServer(m, dyServer, ls.Listener)
		log.Info("Started serveServer")
	}
}

// Shutdown all RPC server currently running.
func (s *Server) Shutdown() {
	for _, ls := range s.lmap {
		ls.Server.Stop()
		ls.Listener.Close()
	}
}

// CleanUp the RPC server.
func (s *Server) CleanUp() {
	for _, ls := range s.lmap {
		if ls != nil {
			ls.Listener.Close()
		}
	}
}

func (s *Server) handleUnaryCall(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	v, err := s.validateMetadata(ctx)
	if err != nil || !v {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	return handler(ctx, req)
}

func (s *Server) handleStreamCall(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	v, err := s.validateMetadata(ss.Context())
	if err != nil || !v {
		return errors.New("Not Valid")
	}
	return handler(srv, ss)
}

func (s *Server) validateMetadata(ctx context.Context) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	var sig []string
	if sig, ok = md["x-signature"]; !ok {
		return false, errors.New("Missing x-signature")
	}
	if len(sig) == 0 {
		return false, errors.New("x-signature array is length 0")
	}
	return s.val.DecodeAndValidate(sig[0]), nil
}

func buildServer(s *grpc.Server, ty string) RPCServer {
	switch ty {
	case "raw":
		return s
	case "web":
		options := []grpcweb.Option{
			grpcweb.WithCorsForRegisteredEndpointsOnly(false),
			grpcweb.WithOriginFunc(func(o string) bool {
				return true
			}),
		}

		wrappedGrpc := grpcweb.WrapServer(s, options...)

		return &WrapperHTTP{
			ser: &http.Server{
				WriteTimeout: 1 * time.Hour,
				ReadTimeout:  10 * time.Second,
				Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
					wrappedGrpc.ServeHTTP(resp, req)
				}),
			},
		}
	}
	return nil
}

func serveServer(ty string, server RPCServer, listener net.Listener) {
	go func() {
		fmt.Printf("RPC server [%s] started on: %s\n", ty, listener.Addr())
		if err := server.Serve(listener); err != nil {
			log.Error(err)
		}
	}()
}
