package main

import (
	"context"
	"fmt"
	"time"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"
	service "github.com/upperwal/go-mesh/interface/service"
	bootstrap "github.com/upperwal/go-mesh/service/bootstrap"
	host "github.com/upperwal/go-mesh/service/host"
	infra "github.com/upperwal/go-mesh/service/infra"
)

// BootstrapPeers is a list of IPFS Bootstrap peers
var BootstrapPeers = []string{
	"/ip4/168.61.157.191/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9",
}

func main() {
	logging.SetLogLevel("application", "DEBUG")
	logging.SetLogLevel("svc-host", "DEBUG")
	logging.SetLogLevel("svc-infra", "DEBUG")
	logging.SetLogLevel("svc-bootstrap", "DEBUG")

	ctx := context.Background()
	app := application.NewApplication(ctx)
	hostSrv := host.NewHostService()
	infraSrv := infra.NewInfraService()
	bootstrapSrv := bootstrap.NewBootstrapService(true, "hgghj", BootstrapPeers)
	fmt.Println(hostSrv.GetName())
	hostSrv.SetPriority(service.PRIORITY_HIGH)

	app.InjectService(hostSrv)
	app.InjectService(infraSrv)
	app.InjectService(bootstrapSrv)
	app.Start()

	time.Sleep(time.Second * 1)
	//app.Stop()

	time.Sleep(time.Second * 30)
}
