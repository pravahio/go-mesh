package pubsub

import (
	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	bc "github.com/upperwal/go-mesh/interface/blockchain"
)

var log = logging.Logger("fpubsub")

type Filter struct {
	r bc.RemoteAccess
}

func NewFilter(ra bc.RemoteAccess) *Filter {
	return &Filter{
		r: ra,
	}
}

func (f *Filter) FilterSubscriber(p peer.ID, t string) bool {
	log.Info("FilterSubscriber")

	if !f.r.IsPeerSubscribed(p, t) {
		log.Info("Dropping msg to:", p)
		return false
	}

	return true
}

func (f *Filter) FilterPublisher(p peer.ID, ts []string) []bool {
	log.Info("FilterPublisher ", p)

	res := make([]bool, len(ts))

	for i, t := range ts {
		if !f.r.IsPeerAPublisher(p, t) {
			log.Info("Dropping msg from:", p)
			res[i] = false
		} else {
			res[i] = true
		}
	}

	return res
}
