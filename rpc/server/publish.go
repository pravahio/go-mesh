package server

import (
	"errors"
	"context"

	rpc "github.com/upperwal/go-mesh/rpc"
)

// RegisterToPublish register a peer and topic with the remote access server.
func (s *Server) RegisterToPublish(ctx context.Context, info *rpc.PeerTopicInfo) (*rpc.Response, error) {
	if s.m.PubService == nil {
		return nil, errors.New("RPC is not running as a publisher")
	}

	err := s.m.PubService.RegisterToPublish(info.GetTopic())
	if err != nil {
		return nil, err
	}

	log.Info("[RPC] Registered to publish on ", info.GetTopic())

	return &rpc.Response{
		Message: "ok",
	}, nil
}

// Publish serves a publish request over RPC
func (s *Server) Publish(ctx context.Context, data *rpc.PublishData) (*rpc.Response, error) {

	if s.m.PubService == nil {
		return nil, errors.New("RPC is not running as a publisher")
	}

	select {
	case <-ctx.Done():
		log.Info("Context Done")
		return &rpc.Response{
			Message: "Context Done",
		}, nil
	default:
	}

	info := data.GetInfo()
	msg := data.GetData()

	log.Info("[RPC] Publishing data on", info.GetTopic())

	err := s.m.PubService.PublishData(info.GetTopic(), msg.GetRaw())

	if err != nil {
		return nil, err
	}

	return &rpc.Response{
		Message: "Done",
	}, nil
}
