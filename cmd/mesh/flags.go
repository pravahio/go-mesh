package main

import (
	cli "github.com/urfave/cli"
)

const (
	REMOTE_ACCESS_URL = "ra"
	BOOTSTRAP_SERVER  = "bs"
	RENDEZVOUS        = "rnz"
	ENABLE_PUB        = "en-pub"
	ENABLE_SUB        = "en-sub"
	DISABLE_RPC       = "dis-raw-rpc"
	ENABLE_WEB_RPC    = "web-rpc"
	DEBUG             = "debug"
	ACCOUNT_FILE      = "account"
	CONFIG_FILE       = "c"
	LISTEN_ADDRESS    = "la"
	RPC_LA            = "rpc-la"
	WEB_RPC_LA        = "web-rpc-la"
	RPC_CERT_PATH     = "rpc-cert-path"
	RPC_KEY_PATH      = "rpc-key-path"
)

var (
	remoteAccess = cli.StringFlag{
		Name:  REMOTE_ACCESS_URL,
		Usage: "remote access url. [http://127.0.0.1:7545]",
	}

	boostrapServer = cli.StringFlag{
		Name:  BOOTSTRAP_SERVER,
		Usage: "bootstrap server multiaddress",
	}

	bootstrapRendezvous = cli.StringFlag{
		Name:  RENDEZVOUS,
		Usage: "set rendezvous point for bootstrapping",
	}

	enablePublisher = cli.BoolFlag{
		Name:  ENABLE_PUB,
		Usage: "starts a publisher node.",
	}

	enableSubscriber = cli.BoolFlag{
		Name:  ENABLE_SUB,
		Usage: "starts a subscriber node.",
	}

	disableRPC = cli.BoolFlag{
		Name:  DISABLE_RPC,
		Usage: "no RPC will be started.",
	}

	enableWebRPC = cli.BoolFlag{
		Name:  ENABLE_WEB_RPC,
		Usage: "enables web based RPC for browsers [http/1.1]",
	}

	enableDebugging = cli.BoolFlag{
		Name:  DEBUG,
		Usage: "enables debug log printing.",
	}

	accountFile = cli.StringFlag{
		Name:  ACCOUNT_FILE,
		Usage: "account file (.msa)",
	}

	configFile = cli.StringFlag{
		Name:  CONFIG_FILE,
		Usage: "config file (.json)",
	}

	lisAdd = cli.StringFlag{
		Name:  LISTEN_ADDRESS,
		Usage: "listen on this address [127.0.0.1:4444]",
	}

	rpcLA = cli.StringFlag{
		Name:  RPC_LA,
		Usage: "RPC will listen on this address [127.0.0.1:4444]",
	}

	webRPCLA = cli.StringFlag{
		Name:  WEB_RPC_LA,
		Usage: "web RPC will listen on this address [127.0.0.1:4445]",
	}

	rpcCertPath = cli.StringFlag{
		Name:  RPC_CERT_PATH,
		Usage: "Certificate [public key] file path for RPC",
	}

	rpcKeyPath = cli.StringFlag{
		Name:  RPC_KEY_PATH,
		Usage: "Key [private key] file path for RPC",
	}
)

func abs() {

}
