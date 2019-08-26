package utils

import (
	"context"
	"fmt"
	mrand "math/rand"
	"strconv"
	"time"

	logging "github.com/ipfs/go-log"
	autonats "github.com/libp2p/go-libp2p-autonat-svc"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/upperwal/go-mesh/application"
	cfg "github.com/upperwal/go-mesh/config"
	bs "github.com/upperwal/go-mesh/service/bootstrap"
	cli "gopkg.in/urfave/cli.v1"
)

const (
	BOOT_RNZ     = "/BOOT_RNZ"
	PORT         = "4000"
	REFRESH_RATE = 15 * time.Minute
)

var (
	ACCOUNT = "a"
	SEED    = "seed"
)

var (
	BootstrapAccountFlag = cli.StringFlag{
		Name:  ACCOUNT,
		Usage: "give an account file (.msa)",
	}
	BootstrapAccountSeedFlag = cli.StringFlag{
		Name:  SEED,
		Usage: "generate account with seed",
	}
)

func BoostrapCommandHandler(ctx *cli.Context) {
	if ctx.Bool(CREATE_ACCOUNT) {
		if s := ctx.String(SEED); s != "" {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println("Incorrect seed value. Should be an int.")
				return
			}
			generateAccountWithSeed(i)
			return
		}
		createAccount(ctx, "boot.msa")
		return
	}
	if f := ctx.String(PARSE); f != "" {
		parse(f)
		return
	}
	if f := ctx.String(ACCOUNT); f != "" {
		m, err := readFromFile(f)
		if err != nil {
			return
		}

		if v, ok := m["Libp2pPrivKey"]; ok {
			k, err := GetLibp2pPrivKey(v)
			if err != nil {
				fmt.Println("Error while reading key from the account file")
				return
			}
			boot(k)
		}
		return
	}

	cli.ShowCommandHelpAndExit(ctx, "bootstrap", 0)
}

func generateAccountWithSeed(s int64) error {
	r := mrand.New(mrand.NewSource(s))
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return err
	}

	m := make(map[string][]byte)

	raw, err := crypto.MarshalPrivateKey(prvKey)
	if err != nil {
		return err
	}

	m["Libp2pPrivKey"] = raw
	writeToFile("boot.msa", m)
	fmt.Println("Writing account data to: boot.msa")

	peerID, err := peer.IDFromPrivateKey(prvKey)
	if err != nil {
		return err
	}

	fmt.Println("Peer ID: ", peerID)

	return nil
}

/* type netNotifiee struct{}

func (nn *netNotifiee) Connected(n inet.Network, c inet.Conn) {
	fmt.Printf("Connected to: %s/p2p/%s\n", c.RemoteMultiaddr(), c.RemotePeer().Pretty())
}

func (nn *netNotifiee) Disconnected(n inet.Network, v inet.Conn)   {}
func (nn *netNotifiee) OpenedStream(n inet.Network, v inet.Stream) {}
func (nn *netNotifiee) ClosedStream(n inet.Network, v inet.Stream) {}
func (nn *netNotifiee) Listen(n inet.Network, a ma.Multiaddr)      {}
func (nn *netNotifiee) ListenClose(n inet.Network, a ma.Multiaddr) {} */

/* func boot(k crypto.PrivKey) {
	logging.SetLogLevel("dht", "DEBUG")
	logging.SetLogLevel("relay", "DEBUG")
	logging.SetLogLevel("pubsub", "DEBUG")
	ctx := context.Background()

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	sourceMultiAddr, _ := ma.NewMultiaddr("/ip4/0.0.0.0/udp/4000/quic")

	tpt, err := libp2pquic.NewTransport(k)
	if err != nil {
		panic(err)
	}
	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(k),
		libp2p.Transport(tpt),
		libp2p.EnableRelay(circuit.OptHop),
	)
	if err != nil {
		panic(err)
	}

	no := &netNotifiee{}
	host.Network().Notify(no)

	fmt.Println("This node: ", host.ID().Pretty(), " ", host.Addrs())

	_, err = dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	//g.Subscribe("GGN.BUS")

	host.SetStreamHandler("BOOTSTRAP", handler)

	select {}
}

func handler(s inet.Stream) {
	fmt.Println("New stream from: ", s.Conn().RemotePeer())
} */

func boot(k crypto.PrivKey) {
	ctx := context.Background()

	logging.SetLogLevel("autonat", "DEBUG")
	logging.SetLogLevel("relay", "DEBUG")
	logging.SetLogLevel("net/identify", "DEBUG")
	logging.SetLogLevel("autonat-svc", "DEBUG")
	fmt.Println("Running Bootstrap node...")

	app, err := application.NewApplication(
		ctx,
		k,
		nil,
		"0.0.0.0",
		PORT,
		true,
	)
	if err != nil {
		log.Warning(err)
		return
	}

	_, err = autonats.NewAutoNATService(ctx, app.GetHost())
	if err != nil {
		log.Error(err)
		return
	}

	bservice := bs.NewBootstrapService(false, BOOT_RNZ, cfg.BootstrapList, REFRESH_RATE)
	app.InjectService(bservice)

	err = app.Start()
	if err != nil {
		log.Warning(err)
		return
	}

	fmt.Println("Waiting for others to connect")

	app.Wait()

}
