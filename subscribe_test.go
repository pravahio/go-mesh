package app_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"

	// Services
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
	subservice "github.com/upperwal/go-mesh/service/subscriber"

	// Driver
	driver "github.com/upperwal/go-mesh/driver"
)

func TestSubscribe(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-subscriber", "DEBUG")
	logging.SetLogLevel("fpubsub", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	bservice := bootservice.NewBootstrapService(false, "abc", []string{"/ip4/127.0.0.1/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9"})
	subService := subservice.NewSubscriberService(driver.NewEthDriver())

	app.InjectService(bservice)
	app.InjectService(subService)

	app.Start()

	time.Sleep(3 * time.Second)
	c, err := subService.SubscribeToTopic("GGN.BUS")
	if err != nil {
		t.Error(err)
	}

	for {
		m := <-c

		fmt.Println(string(m.GetData()))
	}

	app.Wait()
}
