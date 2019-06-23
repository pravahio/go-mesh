package main

import (
	"context"
	"io"
	"net"
	"os"

	logging "github.com/ipfs/go-log"
	rpc "github.com/upperwal/go-mesh/rpc"
	grpc "google.golang.org/grpc"
)

var log = logging.Logger("rpc-server")

type server struct {
	app *Application
}

func (s *server) RegisterToPublish(ctx context.Context, info *rpc.PeerTopicInfo) (*rpc.Response, error) {
	err := s.app.pubService.RegisterToPublish(info.GetTopic())
	if err != nil {
		return nil, err
	}

	log.Info("[RPC] Registered to publish on ", info.GetTopic())

	return &rpc.Response{
		Message: "ok",
	}, nil
}

func (s *server) Publish(stream rpc.Mesh_PublishServer) error {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			log.Info("Done")
			return nil
		default:
		}

		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&rpc.Response{
				Message: "ok",
			})
		}
		if err != nil {
			continue
		}

		info := data.GetInfo()
		msg := data.GetData()

		log.Info("[RPC] Publishing data on", info.GetTopic())

		err = s.app.pubService.PublishData(info.GetTopic(), msg.GetRaw())

		if err != nil {
			log.Warning(err)
		}
	}
}

func (s *server) Subscribe(info *rpc.PeerTopicInfo, stream rpc.Mesh_SubscribeServer) error {
	log.Info("[RPC] Subscribing to ", info.GetTopic())

	ctx := stream.Context()

	msgChan, err := s.app.subService.SubscribeToTopic(info.GetTopic())
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		data := <-msgChan

		err := stream.Send(&rpc.Data{
			Raw: data.GetData(),
		})

		if err != nil {
			log.Warning(err)
		}
	}
}

func main() {
	enableLogging()

	port := "5555"

	log.Info("Starting RPC server on 127.0.0.1:" + port)

	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		os.Exit(0)
	}

	s := grpc.NewServer()

	a, err := NewApplication()
	if err != nil {
		log.Error(err)
	}
	a.AddSubscriber()
	a.AddPublisher()
	a.Start()

	rpc.RegisterMeshServer(s, &server{a})

	if err = s.Serve(lis); err != nil {
		os.Exit(0)
	}
}

func enableLogging() {
	logging.SetLogLevel("rpc-server", "DEBUG")
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-subscriber", "DEBUG")
	logging.SetLogLevel("svc-publisher", "DEBUG")
	logging.SetLogLevel("fpubsub", "DEBUG")
	logging.SetLogLevel("pubsub", "DEBUG")
	logging.SetLogLevel("eth-driver", "DEBUG")
}
