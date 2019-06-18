package blockchain

import peer "github.com/libp2p/go-libp2p-core/peer"

type RemoteAccess interface {
	Subscribe() error
	Publish() error
	IsPeerSubscribed(peer.ID, string) bool
	IsPeerAPublisher(peer.ID, string) bool
}
