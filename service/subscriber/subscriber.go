package subscriber

import (
	"context"

	logging "github.com/ipfs/go-log"
	inet "github.com/libp2p/go-libp2p-core/network"
	bc "github.com/pravahio/go-mesh/interface/ra"
	service "github.com/pravahio/go-mesh/interface/service"
)

var log = logging.Logger("svc-subscriber")

// Name and Version of this service.
const (
	NAME    = "subscriber"
	VERSION = 1
)

type SubscriberService struct {
	service.ApplicationContainer

	ra bc.RemoteAccess

	// Pubsub related
	topicTracker map[string]*TopicWrapper

	// Context
	ctx context.Context
}

func NewSubscriberService(r bc.RemoteAccess) *SubscriberService {
	ss := &SubscriberService{
		ra:           r,
		topicTracker: make(map[string]*TopicWrapper),
	}

	ss.SetNameVersion(NAME, VERSION)

	return ss
}

func (subService *SubscriberService) Start(ctx context.Context) error {
	log.Info(subService.GetName(), "service started")

	subService.ctx = ctx

	log.Info(subService.GetName(), "sub")

	return nil
}

func (subService *SubscriberService) Stop() error {
	return nil
}

func (subService *SubscriberService) Run(stream inet.Stream) {

}

func (subService *SubscriberService) Get(name string) (chan interface{}, error) {
	return nil, nil
}

func (subService *SubscriberService) Set(name string, value interface{}) error {
	return nil
}
