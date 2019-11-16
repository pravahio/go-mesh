package server

import (
	"context"
	"errors"

	rpc "github.com/pravahio/go-mesh/rpc"
)

// Subscribe serves a subscribe request over RPC
func (s *Server) Subscribe(info *rpc.PeerTopicInfo, stream rpc.Mesh_SubscribeServer) error {
	if s.m.SubService == nil {
		return errors.New("RPC is not running as a subscriber")
	}

	log.Info("[RPC] Subscribing to ", info.GetTopics())

	//ctx := stream.Context()

	msgChan, err := s.m.SubService.SubscribeToTopics(info.GetTopics())
	if err != nil {
		log.Error(err)
		return err
	}

	for {
		// TODO: What id the chan is closed.
		data := <-msgChan

		err := stream.Send(&rpc.Data{
			Topic: data.GetTopicIDs(),
			Raw:   data.GetData(),
		})

		if err != nil {
			log.Warning(err)
			return err
		}
		log.Info("RPC Data Sent: ", err)
	}

}

// Unsubscribe serves an unsubscribe request over RPC
func (s *Server) Unsubscribe(ctx context.Context, info *rpc.PeerTopicInfo) (*rpc.Response, error) {
	if s.m.SubService == nil {
		return nil, errors.New("RPC is not running as a subscriber")
	}

	log.Info("[RPC] Unsubscribing to ", info.GetTopics())

	err := s.m.SubService.UnsubscribeToTopics(info.GetTopics())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rpc.Response{
		Message: "ok",
	}, nil
}
