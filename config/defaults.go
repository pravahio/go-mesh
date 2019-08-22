package config

import (
	"strconv"
)

// Alias for true and false
var (
	True  = true
	False = false
)

// DefaultAccountPrvKey defines a nil default account. A new account will be created when the client starts.
var DefaultAccountPrvKey = func(cfg *Config) error {
	cfg.AccountPrvKey = nil
	return nil
}

// DefaultAccountCert defines a nil cert. A new cert will be created when the client starts
var DefaultAccountCert = func(cfg *Config) error {
	cfg.AccountCert = nil
	return nil
}

// DefaultBootstrapNodes assigns few nodes helping in bootstrapping.
var DefaultBootstrapNodes = func(cfg *Config) error {
	cfg.BootstrapNodes = BootstrapList
	return nil
}

// DefaultBootstrapRendezvous is a string to announce and find more nodes related to similar data topic.
// It is unique to each data topic to get better graph neighbours (nodes sending/receiving similar kind of data).
var DefaultBootstrapRendezvous = func(cfg *Config) error {
	cfg.BootstrapRendezvous = ""
	return nil
}

// DefaultRemoteAccessURI defines a remote point URL which is used for authentication of data topics, publishers and subscribers.
// It can be a central server or a blockchain service.
// It should support interface defined in 'interface/blockchain/blockchain.go'
var DefaultRemoteAccessURI = func(cfg *Config) error {
	cfg.RemoteAccessURI = remoteURI
	return nil
}

// DefaultRPCConfig defines listen address and mode of RPC
var DefaultRPCConfig = func(cfg *Config) error {
	cfg.RPC = []RPCConfig{
		RPCConfig{
			Endpoint: RpcURI + strconv.Itoa(RpcPort),
			Mode:     "raw",
		},
	}
	return nil
}

// Defaults contains the conditions for which defaults are defined.
var Defaults = []struct {
	Fallback func(cfg *Config) bool
	Opt      Option
}{
	{
		Fallback: func(cfg *Config) bool { return cfg.AccountPrvKey == nil },
		Opt:      DefaultAccountPrvKey,
	},
	{
		Fallback: func(cfg *Config) bool { return cfg.AccountCert == nil },
		Opt:      DefaultAccountCert,
	},
	{
		Fallback: func(cfg *Config) bool { return cfg.BootstrapNodes == nil },
		Opt:      DefaultBootstrapNodes,
	},
	{
		Fallback: func(cfg *Config) bool { return cfg.BootstrapRendezvous == "" },
		Opt:      DefaultBootstrapRendezvous,
	},
	{
		Fallback: func(cfg *Config) bool { return cfg.RemoteAccessURI == "" },
		Opt:      DefaultRemoteAccessURI,
	},
	{
		Fallback: func(cfg *Config) bool { return cfg.RPC == nil },
		Opt:      DefaultRPCConfig,
	},
}

// DefaultOptions sets some default values.
var DefaultOptions Option = func(cfg *Config) error {
	for _, defaultCfg := range Defaults {
		if !defaultCfg.Fallback(cfg) {
			continue
		}

		cfg.ApplyOpts(defaultCfg.Opt)
	}
	return nil
}
