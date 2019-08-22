package application

import (
	ic "github.com/libp2p/go-libp2p-core/crypto"
	tpt "github.com/libp2p/go-libp2p-core/transport"
	quictpt "github.com/libp2p/go-libp2p-quic-transport"
	//ma "github.com/multiformats/go-multiaddr"
)

var stunServer = "/ip4/40.78.149.141/udp/3000/"

func getQUICTransport(key ic.PrivKey) (tpt.Transport, error) {
	/* stunMA, err := ma.NewMultiaddr(stunServer)
	if err != nil {
		return nil, err
	} */

	/* quicOption := quictpt.TransportOpt{
		EnableStun:  false,
		StunServers: []ma.Multiaddr{stunMA},
	} */

	return quictpt.NewTransport(key)
}
