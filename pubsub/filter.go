package pubsub

import (
	"time"

	cache "github.com/bluele/gcache"
	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	ra "github.com/pravahio/go-mesh/interface/ra"
)

var log = logging.Logger("fpubsub")
var expireTime = 30 * time.Second

type Filter struct {
	r               ra.RemoteAccess
	c               cache.Cache
	isFilterEnabled bool
}

func NewFilter(r ra.RemoteAccess) *Filter {
	if r == nil {
		return &Filter{
			isFilterEnabled: false,
		}
	}
	return &Filter{
		r: r,
		c: cache.New(50).LRU().Build(),
	}
}

func (f *Filter) FilterSubscriber(p peer.ID, t string) bool {
	log.Info("FilterSubscriber")

	if f.isFilterEnabled == false {
		return true
	}

	var v bool

	cValue, err := f.c.Get(p.String() + t)
	if err != nil {
		v, err = f.r.IsPeerSubscribed(p, t)
		if err != nil {
			log.Error(err)
			return false
		}

		err = f.c.SetWithExpire(p.String()+t, v, expireTime)
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

	if f.isFilterEnabled == false {
		for i := range ts {
			res[i] = true
		}
		return res
	}

	for i, t := range ts {

		cValue, err := f.c.Get(p.String() + t)
		if err != nil {
			v, err = f.r.IsPeerAPublisher(p, t)
			if err != nil {
				log.Error(err)
				res[i] = false
				continue
			}

			err = f.c.SetWithExpire(p.String()+t, v, expireTime)
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
