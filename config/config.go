package config

import (
	"crypto/ecdsa"
	"crypto/x509"
	"time"

	logging "github.com/ipfs/go-log"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
)

var log = logging.Logger("config")

// Option defines a function to help assign configs.
type Option func(cfg *Config) error

// Config struct defines what options can be passed to the mesh client.
type Config struct {
	// Accounts
	AccountPrvKey crypto.PrivKey
	AccountCert   *x509.Certificate

	// Bootstrap
	BootstrapNodes       []string
	BootstrapRendezvous  string
	BootstrapRefreshRate time.Duration

	// Remote Access
	RemoteAccessURI        string // TODO: what about the contract address?????
	RemoteAccessPrivateKey *ecdsa.PrivateKey

	// RPC
	RPC []RPCConfig

	// Listening socket
	Host string
	Port string
}

// NewConfig return new configuration with default or new options set.
func NewConfig(opts ...Option) (*Config, error) {
	cfg := &Config{}
	cfg.ApplyOpts(append(opts, DefaultOptions)...)

	return cfg, nil
}

// ApplyOpts executes all option functions.
func (c *Config) ApplyOpts(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}
