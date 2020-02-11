package application

import (
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	floodsub "github.com/libp2p/go-libp2p-pubsub"
	routing "github.com/libp2p/go-libp2p-routing"
)

type libp2pNodeServices struct {
	dht  *dht.IpfsDHT
	psub *floodsub.PubSub
}

func (app *Application) startNode(ctx context.Context, privKey crypto.PrivKey, hostS, portS string, relayService bool) error {

	quicTransport, err := getQUICTransport(privKey)
	if err != nil {
		log.Error(err)
		return err
	}

	relayOpt := []circuit.RelayOpt{
		circuit.OptDiscovery,
	}

	if relayService {
		relayOpt = append(relayOpt, circuit.OptHop)
	}

	dhtFactory := func(h host.Host) (routing.PeerRouting, error) {
		return dht.New(ctx, h)
	}

	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/udp/%s/quic", hostS, portS)),
		//libp2p.NATPortMap(),
		libp2p.Identity(privKey),
		libp2p.Transport(quicTransport),
		libp2p.Routing(dhtFactory),
		libp2p.EnableAutoRelay(),
		libp2p.EnableRelay(relayOpt...),
	)
	if err != nil {
		return err
	}

	log.Info("Multiaddr: ", fmt.Sprintf("/ip4/%s/udp/%s/quic/p2p/%s", hostS, portS, host.ID().Pretty()))

	app.h = host
	app.l = libp2pNodeServices{}

	return nil
}

func (app *Application) startDHT(ctx context.Context) error {
	dhtContext, err := dht.New(ctx, app.h)
	if err != nil {
		return err
	}

	app.l.dht = dhtContext
	return nil
}

func (app *Application) startGossip(ctx context.Context) error {
	gossip, err := floodsub.NewGossipSub(ctx, app.h)
	if err != nil {
		return err
	}

	app.l.psub = gossip
	return nil
}

func (app *Application) SetGossipPeerFilter(f floodsub.PeerFilter) {
	if app.l.psub != nil {
		app.l.psub.SetInOutPeerFilter(f)
	}
}
