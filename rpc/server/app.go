package main

import (
	"context"

	application "github.com/upperwal/go-mesh/application"
	driver "github.com/upperwal/go-mesh/driver/eth"
	ra "github.com/upperwal/go-mesh/interface/blockchain"
	fpubsub "github.com/upperwal/go-mesh/pubsub"
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
	pubservice "github.com/upperwal/go-mesh/service/publisher"
	subservice "github.com/upperwal/go-mesh/service/subscriber"
)

type Application struct {
	app          *application.Application
	remoteAccess ra.RemoteAccess
	subService   *subservice.SubscriberService
	pubService   *pubservice.PublisherService
}

func NewApplication() (*Application, error) {

	app, err := application.NewApplication(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}

	ethDriver, err := driver.NewEthDriver()
	if err != nil {
		return nil, err
	}

	bservice := bootservice.NewBootstrapService(false, "abc", []string{"/ip4/127.0.0.1/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9"})
	app.InjectService(bservice)

	f := fpubsub.NewFilter(ethDriver)
	app.SetGossipPeerFilter(f)

	return &Application{
		app:          app,
		remoteAccess: ethDriver,
	}, nil

	/* app.InjectService(subService)



	app.Start()

	time.Sleep(3 * time.Second)
	c, err := subService.SubscribeToTopic("GGN.BUS")
	if err != nil {
		log.Error(err)
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

	app.Wait() */
}

func (app *Application) AddSubscriber() {
	subService := subservice.NewSubscriberService(app.remoteAccess)
	app.app.InjectService(subService)

	app.subService = subService
}

func (app *Application) AddPublisher() {
	pubService := pubservice.NewPublisherService(app.remoteAccess)
	app.app.InjectService(pubService)

	app.pubService = pubService
}

func (app *Application) Start() error {
	return app.app.Start()
}
