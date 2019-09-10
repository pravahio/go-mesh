package main

import (
	utils "github.com/pravahio/go-mesh/cmd/utils"
	cli "github.com/urfave/cli"
)

var (
	accountCommand = cli.Command{
		Name:   "account",
		Action: utils.AccountCommandHandler,
		Usage:  "handles all account related functions",
		Flags: []cli.Flag{
			utils.CreateAccountFlag,
			utils.AccountOutputFileNameFlag,
			utils.PrintPublicKeyFlag,
			utils.ParseAccountFile,
		},
	}

	boostrapCommand = cli.Command{
		Name:   "bootstrap",
		Action: utils.BoostrapCommandHandler,
		Usage:  "handles bootstrap related functions",
		Flags: []cli.Flag{
			utils.CreateAccountFlag,
			utils.AccountOutputFileNameFlag,
			utils.PrintPublicKeyFlag,
			utils.ParseAccountFile,
			utils.BootstrapAccountFlag,
			utils.BootstrapAccountSeedFlag,
		},
	}
)
