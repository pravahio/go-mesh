package app_test

import (
	"context"
	"testing"
	"time"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"

	// Services
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
	subservice "github.com/upperwal/go-mesh/service/subscriber"

	// Driver
	driver "github.com/upperwal/go-mesh/driver/eth"

	cfg "github.com/upperwal/go-mesh/config"
)

func TestApp(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-subscriber", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil, "0.0.0.0", "0", false)
	if err != nil {
		t.Fatal(err)
	}

	s, err := driver.NewEthDriver("", nil)
	if err != nil {
		t.Fatal(err)
	}

	bservice := bootservice.NewBootstrapService(
		false,
		"abc",
		cfg.BootstrapList,
		15*time.Second)
	subService := subservice.NewSubscriberService(s)

	app.InjectService(bservice)
	app.InjectService(subService)

	err = app.Start()
	if err != nil {
		t.Fatal(err)
	}

}
