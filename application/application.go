package application

import (
	"context"
	"crypto/x509"
	"sync"

	logging "github.com/ipfs/go-log"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"

	service "github.com/upperwal/go-mesh/interface/service"
	account "github.com/upperwal/go-mesh/util/account"
)

var log = logging.Logger("application")

type Application struct {
	// Libp2p Node/Host related
	h host.Host
	l libp2pNodeServices

	// Service related
	ServiceMapMutex *sync.Mutex
	ServiceMap      map[string]service.ServiceData

	// Application contexts
	ctx           context.Context
	ctxCancelFunc context.CancelFunc
}

func NewApplication(ctx context.Context, privKey crypto.PrivKey, cert *x509.Certificate) (*Application, error) {

	appCtx, appCancel := context.WithCancel(ctx)

	a := &Application{
		ServiceMapMutex: &sync.Mutex{},
		ServiceMap:      make(map[string]service.ServiceData),
		ctx:             appCtx,
		ctxCancelFunc:   appCancel,
	}

	var err error
	if privKey == nil {
		cert, privKey, _, err = account.GenerateAccount()
		if err != nil {
			log.Info(err)
			return nil, err
		}
		log.Info("New account generated")
	}

	err = a.startNode(a.ctx, privKey)
	if err != nil {
		return nil, err
	}
	err = a.startDHT(a.ctx)
	if err != nil {
		return nil, err
	}
	err = a.startGossip(a.ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (app *Application) Start() error {
	log.Info("Starting application")
	for _, srv := range app.ServiceMap {
		go func(srv service.Service) {
			err := srv.Start(app.ctx)
			if err != nil {
				log.Error("Error while Starting service:\x1B[35m", srv.GetName(), "\x1B[0m=> Service exiting with error: \x1B[31m", err, "\x1B[0m")
				return
			}
		}(srv.(service.Service))
	}
	return nil
}

func (app *Application) Stop() error {
	app.ctxCancelFunc()
	return nil
}

func (app *Application) GetService(name string) service.ServiceData {
	return app.ServiceMap[name]
}

func (app *Application) InjectService(srv service.Service) {
	name := srv.(service.Service).GetName()
	protocol := srv.(service.Service).GetProtocol()

	app.ServiceMapMutex.Lock()
	app.ServiceMap[name] = srv
	app.ServiceMapMutex.Unlock()

	app.h.SetStreamHandler(protocol, srv.Run)
	srv.SetAppMeta(app.GetService, &app.h, app.l.dht, app.l.psub)
}

func (app *Application) Wait() {
	<-app.ctx.Done()
}
