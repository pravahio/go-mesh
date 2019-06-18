package app_test

import (
	"context"
	"testing"
	"time"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"
	driver "github.com/upperwal/go-mesh/driver"
	fpubsub "github.com/upperwal/go-mesh/pubsub"
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
	pubservice "github.com/upperwal/go-mesh/service/publisher"
)

func TestPublish(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-publisher", "DEBUG")
	logging.SetLogLevel("fpubsub", "DEBUG")
	logging.SetLogLevel("pubsub", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	ethDriver := driver.NewEthDriver()

	bservice := bootservice.NewBootstrapService(false, "abc", []string{"/ip4/127.0.0.1/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9"})
	pubService := pubservice.NewPublisherService(ethDriver)

	app.InjectService(bservice)
	app.InjectService(pubService)

	f := fpubsub.NewFilter(ethDriver)
	app.SetGossipPeerFilter(f)

	app.Start()

	time.Sleep(3 * time.Second)

	pubService.RegisterToPublish("GGN.BUS")
	for {
		pubService.PublishData("GGN.BUS", []byte("Zoka"))

		time.Sleep(500 * time.Millisecond)
	}
	app.Wait()
}
