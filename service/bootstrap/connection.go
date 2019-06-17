package bootstrap

import (
	addr "github.com/ipfs/go-ipfs-addr"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

func (bs *BootstrapService) startBootstrapping() error {
	/* v1b := cid.V1Builder{
		Codec:  cid.Raw,
		MhType: mh.SHA2_256,
	}

	rendezvousPoint, _ := v1b.Sum([]byte(bs.rendezvousPoint)) */

	return bs.connectToBootstrapPeers()
}

func (bs *BootstrapService) connectToBootstrapPeers() error {
	host := bs.GetHost()
	for _, peerAddr := range bs.bootstrapPeers {
		ipfsAddr, err := addr.ParseString(peerAddr)
		if err != nil {
			return err
		}

		peerInfo, err := pstore.InfoFromP2pAddr(ipfsAddr.Multiaddr())
		if err != nil {
			return err
		}

		if err := host.Connect(bs.ctxLocal, *peerInfo); err != nil {
			log.Error("Could not connect to bootstrap peer: ", peerAddr, err)
			return err
		} else {
			log.Info("Connectivity with bootstrap peer success! Peer: ", peerAddr)
		}

		bs.s, err = host.NewStream(bs.ctxLocal, peerInfo.ID, "BOOTSTRAP")
		if err != nil {
			return err
		}
		log.Info("Started new bootstrap stream")
	}

	return nil
}
