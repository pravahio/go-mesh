package app_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"
	driver "github.com/upperwal/go-mesh/driver/eth"
	fpubsub "github.com/upperwal/go-mesh/pubsub"
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
	subservice "github.com/upperwal/go-mesh/service/subscriber"
)

func TestSubscribe(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-subscriber", "DEBUG")
	logging.SetLogLevel("fpubsub", "DEBUG")
	logging.SetLogLevel("pubsub", "DEBUG")
	logging.SetLogLevel("eth-driver", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	ethDriver, err := driver.NewEthDriver()
	if err != nil {
		t.Fatal(err)
	}

	bservice := bootservice.NewBootstrapService(false, "abc", []string{"/ip4/127.0.0.1/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9"})
	subService := subservice.NewSubscriberService(ethDriver)

	app.InjectService(bservice)
	app.InjectService(subService)

	f := fpubsub.NewFilter(ethDriver)
	app.SetGossipPeerFilter(f)

	app.Start()

	time.Sleep(3 * time.Second)
	c, err := subService.SubscribeToTopic("GGN.BUS")
	if err != nil {
		t.Fatal(err)
	}

	counter := 0
	tm := time.Now()
	for {
		<-c
		counter++
		//fmt.Println(string(m.GetData()))
		if counter%100 == 0 {
			fmt.Println("Messages Recv Rate / Count: ", float64(counter)/time.Since(tm).Seconds(), counter)
		}
	}

	app.Wait()
}
