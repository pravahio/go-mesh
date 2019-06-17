package app_test

import (
	"context"
	"testing"

	logging "github.com/ipfs/go-log"
	application "github.com/upperwal/go-mesh/application"

	// Services
	bootservice "github.com/upperwal/go-mesh/service/bootstrap"
)

func TestApp(t *testing.T) {
	logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG")

	app, err := application.NewApplication(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	bservice := bootservice.NewBootstrapService(false, "abc", []string{"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"})

	app.InjectService(bservice)

	app.Start()
}
