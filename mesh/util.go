package mesh

import (
	"errors"

	config "github.com/pravahio/go-mesh/config"
)

func checkConfig(cfg *config.Config) error {
	if cfg.BootstrapRendezvous == "" {
		return errors.New("BootstrapRendezvous is absolutely needed based on the data topic")
	}
	return nil
}
