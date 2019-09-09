package app_test

import (
	"context"
	"testing"
	"time"

	logging "github.com/ipfs/go-log"
	application "github.com/pravahio/go-mesh/application"
	cfg "github.com/pravahio/go-mesh/config"
	driver "github.com/pravahio/go-mesh/driver/eth"
	fpubsub "github.com/pravahio/go-mesh/pubsub"
	bootservice "github.com/pravahio/go-mesh/service/bootstrap"
	subservice "github.com/pravahio/go-mesh/service/subscriber"
)

func TestSubscribe(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-subscriber", "DEBUG")
	logging.SetLogLevel("fpubsub", "DEBUG")
	logging.SetLogLevel("pubsub", "DEBUG")
	logging.SetLogLevel("eth-driver", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil, "0.0.0.0", "0", false)
	if err != nil {
		t.Fatal(err)
	}

	ethDriver, err := driver.NewEthDriver("", nil)
	if err != nil {
		t.Fatal(err)
	}

	bservice := bootservice.NewBootstrapService(
		false,
		"abc",
		cfg.BootstrapList,
		15*time.Second)
	subService := subservice.NewSubscriberService(ethDriver)

	app.InjectService(bservice)
	app.InjectService(subService)

	f := fpubsub.NewFilter(ethDriver)
	app.SetGossipPeerFilter(f)

	err = app.Start()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(3 * time.Second)
	_, err = subService.SubscribeToTopic("GGN.BUS")
	if err != nil {
		t.Fatal(err)
	}
}
