package pubsub

import (
	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

var log = logging.Logger("fpubsub")

type Filter struct{}

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) InFilter(p peer.ID) bool {
	log.Info("Inward peer: ", p)
	return true
}

func (f *Filter) OutFilter(p peer.ID) bool {
	log.Info("Outward peer: ", p)
	return true
}
