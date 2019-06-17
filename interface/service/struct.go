package service

import (
	"fmt"

	host "github.com/libp2p/go-libp2p-core/host"
	proto "github.com/libp2p/go-libp2p-core/protocol"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	floodsub "github.com/libp2p/go-libp2p-pubsub"
)

type ApplicationCallback func(string) ServiceData

type ApplicationContainer struct {
	GetService ApplicationCallback
	Priority   int // higher the better
	h          *host.Host
	d          *dht.IpfsDHT
	p          *floodsub.PubSub

	// Meta data
	name     string
	version  uint
	protocol proto.ID
}

func (sas *ApplicationContainer) SetAppMeta(app ApplicationCallback, h *host.Host, d *dht.IpfsDHT, p *floodsub.PubSub) {
	sas.GetService = app
	sas.h = h
	sas.d = d
	sas.p = p
}

func (sas *ApplicationContainer) SetNameVersion(name string, ver uint) {
	sas.name = name
	sas.version = ver
	sas.protocol = proto.ID("/" + name + "/v" + fmt.Sprintf("%d", ver))
}

func (sas *ApplicationContainer) SetPriority(p int) {
	sas.Priority = p
}

func (sas *ApplicationContainer) GetHost() host.Host {
	return *sas.h
}

func (sas *ApplicationContainer) GetDHT() *dht.IpfsDHT {
	return sas.d
}

func (sas *ApplicationContainer) GetPubSub() *floodsub.PubSub {
	return sas.p
}

func (sas *ApplicationContainer) GetName() string {
	return sas.name
}

func (sas *ApplicationContainer) GetVersion() uint {
	return sas.version
}

func (sas *ApplicationContainer) GetProtocol() proto.ID {
	return sas.protocol
}
