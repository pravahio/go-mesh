package application

import (
	"context"

	libp2p "github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	floodsub "github.com/libp2p/go-libp2p-pubsub"
)

type libp2pNodeServices struct {
	dht  *dht.IpfsDHT
	psub *floodsub.PubSub
}

func (app *Application) startNode(ctx context.Context, privKey crypto.PrivKey) error {

	quicTransport, err := getQUICTransport(privKey)
	if err != nil {
		log.Error(err)
		return err
	}

	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/udp/0/quic"),
		libp2p.NATPortMap(),
		libp2p.Identity(privKey),
		libp2p.Transport(quicTransport),
		libp2p.EnableRelay(circuit.OptDiscovery),
	)
	if err != nil {
		return err
	}

	log.Info("Host ID: ", host.ID().Pretty())

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
