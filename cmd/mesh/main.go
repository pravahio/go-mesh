package main

import (
	"context"
	"os"
	"strconv"

	logging "github.com/ipfs/go-log"
	"github.com/upperwal/go-mesh/config"
	mclient "github.com/upperwal/go-mesh/mesh"
	rpc "github.com/upperwal/go-mesh/rpc/server"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	app = cli.NewApp()
	log = logging.Logger("mesh-cli")
)

func init() {
	app.Name = "mesh"
	app.Version = "v0.0.1"
	app.Usage = "go-mesh command line interface"
	app.Description = "go-mesh command line interface"
	app.Authors = []cli.Author{
		{Name: "Abhishek Upperwal", Email: "abhishek@soket.in"},
	}
	app.Action = mesh
	app.Flags = []cli.Flag{
		remoteAccess,
		boostrapServer,
		bootstrapRendezvous,
		enablePublisher,
		enableSubscriber,
		disableRPC,
		enableWebRPC,
		enableDebugging,
		accountFile,
		configFile,
	}
}

func mesh(ctx *cli.Context) {

	c := NewConfig(ctx)

	applyLogs(c.Bool(DEBUG))

	m, err := mclient.NewMesh(
		context.Background(),
		getBootRendz(c.String(RENDEZVOUS)),
		getBootServer(c.String(BOOTSTRAP_SERVER)),
		applyRA(c.String(REMOTE_ACCESS_URL)),
	)
	if err != nil {
		log.Fatal(err)
	}

	applyNodeType(
		c.Bool(ENABLE_PUB),
		c.Bool(ENABLE_SUB),
		m,
	)

	applyRPC(m.Cfg, c.Bool(ENABLE_WEB_RPC), c.Bool(DISABLE_RPC))
	rpcs, err := rpc.NewServer(m)
	if err != nil {
		log.Fatal(err)
	}

	m.Start()
	rpcs.Start()

	m.Wait()
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func getBootRendz(b string) config.Option {
	if b == "" {
		log.Error("Need a rendezvous string")
		os.Exit(1)
	}
	return func(cfg *config.Config) error {
		cfg.BootstrapRendezvous = b
		return nil
	}
}

func getBootServer(b string) config.Option {
	return func(cfg *config.Config) error {
		cfg.BootstrapNodes = []string{
			b,
		}
		return nil
	}
}

func applyRA(s string) config.Option {
	return func(cfg *config.Config) error {
		cfg.RemoteAccessURI = s
		return nil
	}
}

func applyRPC(m *config.Config, en_web, dis_raw bool) {
	if dis_raw {
		m.RPC = []config.RPCConfig{}
	}

	if en_web {
		m.RPC = append(m.RPC, config.RPCConfig{
			Endpoint: config.RpcURI + strconv.Itoa(config.RpcPort+1),
			Mode:     "web",
		})
	}
}

func applyNodeType(p bool, s bool, m *mclient.Mesh) {
	if p {
		log.Info("Configured as a publisher node.")
		m.AddPublisher()
	} else if s {
		log.Info("Configured as a subscriber node.")
		m.AddSubscriber()
	}
}

func applyLogs(b bool) {
	if b {
		logging.SetLogLevel("mesh-cli", "DEBUG")
		logging.SetLogLevel("rpc-server", "DEBUG")
		logging.SetLogLevel("application", "DEBUG")
		logging.SetLogLevel("svc-bootstrap", "DEBUG")
		logging.SetLogLevel("application", "DEBUG")
		logging.SetLogLevel("svc-publisher", "DEBUG")
		logging.SetLogLevel("fpubsub", "DEBUG")
		logging.SetLogLevel("pubsub", "DEBUG")
		logging.SetLogLevel("eth-driver", "DEBUG")
	}
}
