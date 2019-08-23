package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	grpcweb "github.com/improbable-eng/grpc-web/go/grpcweb"
	logging "github.com/ipfs/go-log"
	config "github.com/upperwal/go-mesh/config"
	mesh "github.com/upperwal/go-mesh/mesh"
	rpc "github.com/upperwal/go-mesh/rpc"
	grpc "google.golang.org/grpc"
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
	m    *mesh.Mesh
	lmap map[string]*LisServer
}

// NewServer creates a new RPC server.
func NewServer(m *mesh.Mesh) (*Server, error) {

	s := Server{
		m:    m,
		lmap: make(map[string]*LisServer),
	}

	err := s.registerRPC(m.Cfg.RPC)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Server) registerRPC(rpcConfig []config.RPCConfig) error {

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

func (s *Server) Start() {
	grpcServer := grpc.NewServer()
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
