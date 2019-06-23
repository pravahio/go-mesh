package blockchain

import peer "github.com/libp2p/go-libp2p-core/peer"

type RemoteAccess interface {
	Subscribe(peer.ID, string) error
	Publish(peer.ID, string) error
	IsPeerSubscribed(peer.ID, string) (bool, error)
	IsPeerAPublisher(peer.ID, string) (bool, error)
}
