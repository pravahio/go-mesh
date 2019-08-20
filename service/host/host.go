package host

import (
	"context"
	"errors"

	logging "github.com/ipfs/go-log"
	inet "github.com/libp2p/go-libp2p-core/network"
	service "github.com/upperwal/go-mesh/interface/service"
	datatype "github.com/upperwal/go-mesh/service/host/datatype"
)

var log = logging.Logger("svc-host")

const (
	NAME    = "host"
	VERSION = 1
)

// HostService represents a host service.
type HostService struct {
	service.ApplicationContainer

	// Data Container
	loc datatype.Location

	// Context related
	ctxLocal           context.Context
	ctxLocalCancelFunc context.CancelFunc
}

// NewHostService creates a new HostService.
func NewHostService() *HostService {
	hs := &HostService{
		loc: datatype.Location{},
	}

	hs.SetNameVersion(NAME, VERSION)
	return hs
}

func (hs *HostService) Start(ctx context.Context) error {
	log.Info("host service started")
	hs.ctxLocal, hs.ctxLocalCancelFunc = context.WithCancel(ctx)

	bootstrapSvc := hs.GetService("bootstrap")
	if bootstrapSvc == nil {
		return errors.New("bootstrap service is missing")
	}
	r, err := bootstrapSvc.Get("ready")
	if err != nil {
		return err
	}

	ready := (<-r).(bool)
	if !ready {
		return errors.New("Something is wrong with bootstrap service")
	}

	return nil
}

func (hs *HostService) Stop() error {
	hs.ctxLocalCancelFunc()
	return nil
}

func (hs *HostService) Run(s inet.Stream) {
	for {
		log.Info("Running host")
		select {
		case <-hs.ctxLocal.Done():
			log.Info("Context Expired")
			return
		}
	}
	return
}

// Get implements service.ServiceData
func (hs *HostService) Get(key string) (chan interface{}, error) {
	switch key {
	case "location":
		loc := make(chan interface{}, 1)
		loc <- hs.loc
		return loc, nil
	}
	return nil, errors.New("key unknown")
}

// Set implements service.ServiceData
func (hs *HostService) Set(key string, value interface{}) error {
	return service.ERR_PERMISSION_DENIED
}

func (hs *HostService) UpdateLocation(longitude, latitude float64) {
	hs.loc.Longitude = longitude
	hs.loc.Latitude = latitude
}
