package app_test

import (
	"context"
	"testing"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"

	// Services
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
	subservice "github.com/upperwal/go-mesh/service/subscriber"

	// Driver
	driver "github.com/upperwal/go-mesh/driver/eth"
)

func TestApp(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-subscriber", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	s, err := driver.NewEthDriver()
	if err != nil {
		t.Fatal(err)
	}

	bservice := bootservice.NewBootstrapService(false, "abc", []string{"/ip4/127.0.0.1/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9"})
	subService := subservice.NewSubscriberService(s)

	app.InjectService(bservice)
	app.InjectService(subService)

	app.Start()
	app.Wait()
}
