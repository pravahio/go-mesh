package service

import (
	"context"

	host "github.com/libp2p/go-libp2p-core/host"
	inet "github.com/libp2p/go-libp2p-core/network"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	floodsub "github.com/libp2p/go-libp2p-pubsub"
)

// ServiceData represents functions when GetService
// is called in any service.
type ServiceData interface {
	Get(string) (chan interface{}, error)
	Set(string, interface{}) error
}

// ApplicationLinker is implemented in struct
type ApplicationLinker interface {
	SetAppMeta(ApplicationCallback, *host.Host, *dht.IpfsDHT, *floodsub.PubSub)
}

// Service defines a generic service.
type Service interface {
	GetName() string
	GetVersion() uint
	GetProtocol() protocol.ID
	Start(context.Context) error
	Stop() error
	Run(inet.Stream)
}
