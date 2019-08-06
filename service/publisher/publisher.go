package publisher

import (
	"context"

	logging "github.com/ipfs/go-log"
	inet "github.com/libp2p/go-libp2p-core/network"
	ra "github.com/upperwal/go-mesh/interface/ra"
	service "github.com/upperwal/go-mesh/interface/service"
)

var log = logging.Logger("svc-publisher")

// Name and Version of this service.
const (
	NAME    = "publisher"
	VERSION = 1
)

type TopicName string

type PublisherService struct {
	service.ApplicationContainer

	ra ra.RemoteAccess

	// Pubsub related
	//topicTracker map[TopicName]*TopicWrapper

	// Context
	ctx context.Context
}

func NewPublisherService(r ra.RemoteAccess) *PublisherService {
	ps := &PublisherService{
		ra: r,
	}

	ps.SetNameVersion(NAME, VERSION)

	return ps
}

func (pubService *PublisherService) Start(ctx context.Context) error {
	log.Info(pubService.GetName(), "service started")

	pubService.ctx = ctx

	log.Info(pubService.GetName(), "pub")

	return nil
}

func (pubService *PublisherService) Stop() error {
	return nil
}

func (pubService *PublisherService) Run(stream inet.Stream) {

}

func (pubService *PublisherService) Get(name string) (chan interface{}, error) {
	return nil, nil
}

func (pubService *PublisherService) Set(name string, value interface{}) error {
	return nil
}
