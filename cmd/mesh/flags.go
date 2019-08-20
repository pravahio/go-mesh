package main

import (
	cli "gopkg.in/urfave/cli.v1"
)

const (
	REMOTE_ACCESS_URL = "ra"
	BOOTSTRAP_SERVER  = "bs"
	RENDEZVOUS        = "rnz"
	ENABLE_PUB        = "en_pub"
	ENABLE_SUB        = "en_sub"
	DISABLE_RPC       = "dis_raw_rpc"
	ENABLE_WEB_RPC    = "web_rpc"
	DEBUG             = "debug"
	ACCOUNT           = "account"
	CONFIG_FILE       = "config_file"
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
		Name:  ACCOUNT,
		Usage: "account file (.msa)",
	}

	configFile = cli.StringFlag{
		Name:  CONFIG_FILE,
		Usage: "config file (.json)",
	}
)

func abs() {

}
