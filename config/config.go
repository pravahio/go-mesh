package config

import (
	"crypto/x509"

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
	BootstrapNodes      []string
	BootstrapRendezvous string

	// Remote Access
	RemoteAccessURI string
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
