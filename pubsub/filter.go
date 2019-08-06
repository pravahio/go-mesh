package pubsub

import (
	cache "github.com/bluele/gcache"
	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	ra "github.com/upperwal/go-mesh/interface/ra"
)

var log = logging.Logger("fpubsub")

type Filter struct {
	r ra.RemoteAccess
	c cache.Cache
}

func NewFilter(r ra.RemoteAccess) *Filter {
	return &Filter{
		r: r,
		c: cache.New(50).LRU().Build(),
	}
}

func (f *Filter) FilterSubscriber(p peer.ID, t string) bool {
	log.Info("FilterSubscriber")

	var v bool

	cValue, err := f.c.Get(p.String() + t)
	if err != nil {
		v, err = f.r.IsPeerSubscribed(p, t)
		if err != nil {
			log.Error("IsPeerSubscribed returned err")
			return false
		}

		err = f.c.Set(p.String()+t, v)
		if err != nil {
			log.Error(err)
		}
	} else {
		v = cValue.(bool)
	}

	if !v {
		log.Info("Dropping msg to:", p)
	}

	return v
}

func (f *Filter) FilterPublisher(p peer.ID, ts []string) []bool {
	log.Info("FilterPublisher ", p)

	var v bool
	res := make([]bool, len(ts))

	for i, t := range ts {

		cValue, err := f.c.Get(p.String() + t)
		if err != nil {
			v, err = f.r.IsPeerAPublisher(p, t)
			if err != nil {
				log.Error("IsPeerAPublisher returned err")
				res[i] = false
				continue
			}

			err = f.c.Set(p.String()+t, v)
			if err != nil {
				log.Error(err)
			}
		} else {
			v = cValue.(bool)
		}

		if !v {
			log.Info("Dropping msg from:", p, "for topic", t)
		}

		res[i] = v
	}

	return res
}
