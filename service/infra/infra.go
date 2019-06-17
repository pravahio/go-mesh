package infra

import (
	"context"

	logging "github.com/ipfs/go-log"
	inet "github.com/libp2p/go-libp2p-core/network"
	service "github.com/upperwal/go-mesh/interface/service"
)

var log = logging.Logger("svc-infra")

const (
	NAME    = "infra"
	VERSION = 1
)

type InfraService struct {
	service.ApplicationContainer
	name    string
	version uint

	// Context related
	ctxLocal           context.Context
	ctxLocalCancelFunc context.CancelFunc
}

func NewInfraService() *InfraService {
	is := &InfraService{}

	is.SetNameVersion(NAME, VERSION)
	return is
}

func (is *InfraService) Start(ctx context.Context) error {
	log.Info("infra service started")
	is.ctxLocal, is.ctxLocalCancelFunc = context.WithCancel(ctx)
	return nil
}

func (is *InfraService) Stop() error {
	is.ctxLocalCancelFunc()
	return nil
}

func (is *InfraService) Run(inet.Stream) {
	return
}

// Get implements service.ServiceData
func (is *InfraService) Get(key string) (chan interface{}, error) {
	return nil, nil
}

// Set implements service.ServiceData
func (is *InfraService) Set(key string, value interface{}) error {
	return service.ERR_PERMISSION_DENIED
}

func (hs *InfraService) Some() {
	/* srv := hs.GetService("host")
	srv.Get() */
}
