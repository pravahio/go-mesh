package bootstrap

import (
	"context"
	"fmt"
	"time"

	logging "github.com/ipfs/go-log"
	inet "github.com/libp2p/go-libp2p-core/network"
	service "github.com/pravahio/go-mesh/interface/service"
)

var log = logging.Logger("svc-bootstrap")

// Name and Version of this service.
const (
	NAME    = "bootstrap"
	VERSION = 1
)

// HostService represents a host service.
type BootstrapService struct {
	service.ApplicationContainer

	// Data container
	announce        bool
	rendezvousPoint string
	bootstrapPeers  []string

	// ready data
	ready         bool
	readyChanList []chan interface{}

	// Context related
	ctxLocal           context.Context
	ctxLocalCancelFunc context.CancelFunc

	// refresh rate
	d time.Duration

	s inet.Stream
}

// NewBootstrapService creates a new BootstrapService.
func NewBootstrapService(ann bool, rPoint string, bList []string, d time.Duration) *BootstrapService {
	hs := &BootstrapService{
		announce:        ann,
		rendezvousPoint: rPoint,
		bootstrapPeers:  bList,
		readyChanList:   make([]chan interface{}, 0),
		d:               d,
	}

	hs.SetNameVersion(NAME, VERSION)
	return hs
}

func (bs *BootstrapService) Start(ctx context.Context) error {
	log.Info(bs.GetName(), "service started")
	bs.ctxLocal, bs.ctxLocalCancelFunc = context.WithCancel(ctx)

	err := bs.startBootstrapping()
	if err != nil {

		bs.ready = false
		for _, rChan := range bs.readyChanList {
			rChan <- false
		}

		return err
	}

	bs.ready = true
	for _, rChan := range bs.readyChanList {
		rChan <- true
	}

	return nil
}

func (bs *BootstrapService) Stop() error {
	bs.ctxLocalCancelFunc()
	bs.GetHost().Network().Close()
	return nil
}

func (bs *BootstrapService) Run(s inet.Stream) {
	fmt.Println("My info:", bs.GetHost().ID(), "(", bs.GetHost().Addrs(), ") => Remote:", s.Conn().RemotePeer(), s.Conn().RemoteMultiaddr())
}

// Get implements service.ServiceData
func (bs *BootstrapService) Get(key string) (chan interface{}, error) {
	switch key {
	case "ready":
		newReadyChan := make(chan interface{}, 1)
		if bs.ready {
			newReadyChan <- bs.ready
		} else {
			bs.readyChanList = append(bs.readyChanList, newReadyChan)
		}

		return newReadyChan, nil
	}
	return nil, service.ERR_DEFAULT
}

// Set implements service.ServiceData
func (bs *BootstrapService) Set(key string, value interface{}) error {
	return service.ERR_PERMISSION_DENIED
}

func (bs *BootstrapService) WriteTemp() error {
	n, err := bs.s.Write([]byte("hello"))
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("Written ", n, "bytes")
	return nil
}
