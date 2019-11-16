package server

import (
	"context"
	"errors"

	rpc "github.com/pravahio/go-mesh/rpc"
)

// RegisterToPublish register a peer and topic with the remote access server.
func (s *Server) RegisterToPublish(ctx context.Context, info *rpc.PeerTopicInfo) (*rpc.Response, error) {
	if s.m.PubService == nil {
		return nil, errors.New("RPC is not running as a publisher")
	}

	err := s.m.PubService.RegisterToPublish(info.GetTopics())
	if err != nil {
		return nil, err
	}

	log.Info("[RPC] Registered to publish on ", info.GetTopics())

	return &rpc.Response{
		Message: "ok",
	}, nil
}

// UnregisterToPublish unregister a peer and topic with the remote access server.
func (s *Server) UnregisterToPublish(ctx context.Context, info *rpc.PeerTopicInfo) (*rpc.Response, error) {
	if s.m.PubService == nil {
		return nil, errors.New("RPC is not running as a publisher")
	}

	err := s.m.PubService.UnregisterToPublish(info.GetTopics())
	if err != nil {
		return nil, err
	}

	log.Info("[RPC] Unregistered to publish on ", info.GetTopics())

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

	log.Info("[RPC] Publishing data on", info.GetTopics())

	err := s.m.PubService.PublishData(msg.GetRaw(), info.GetTopics())

	if err != nil {
		return nil, err
	}

	return &rpc.Response{
		Message: "Done",
	}, nil
}
