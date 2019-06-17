 package subscriber

import (
	"context"

	inet "github.com/libp2p/go-libp2p-core/network"
	service "github.com/upperwal/go-mesh/interface/service"
)

type SubscriberService struct {
	service.ApplicationContainer
}

func NewSubscriberService() *SubscriberService {
	return &SubscriberService{}
}

func (subService *SubscriberService) Start(ctx context.Context) error {
	return nil
}

func (subService *SubscriberService) Stop() error {
	return nil
}

func (subService *SubscriberService) Run(stream inet.Stream) {

}
