package eth_test

import (
	"testing"

	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	eth "github.com/upperwal/go-mesh/driver/eth"
)

func TestEthSub(t *testing.T) {
	logging.SetLogLevel("eth-driver", "DEBUG")

	d, err := eth.NewEthDriver()
	if err != nil {
		t.Error(err)
	}

	id, err := peer.IDB58Decode("QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9")
	if err != nil {
		t.Error(err)
	}

	err = d.Subscribe(id, "topic")
	if err != nil {
		t.Error(err)
	}
}

func TestEthPub(t *testing.T) {
	logging.SetLogLevel("eth-driver", "DEBUG")

	d, err := eth.NewEthDriver()
	if err != nil {
		t.Error(err)
	}

	id, err := peer.IDB58Decode("QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9")
	if err != nil {
		t.Error(err)
	}

	err = d.Publish(id, "topic")
	if err != nil {
		t.Error(err)
	}
}

func TestEthIsSub(t *testing.T) {
	logging.SetLogLevel("eth-driver", "DEBUG")

	d, err := eth.NewEthDriver()
	if err != nil {
		t.Error(err)
	}

	id, err := peer.IDB58Decode("QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9")
	if err != nil {
		t.Error(err)
	}

	b, err := d.IsPeerSubscribed(id, "t")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("IsPeerSubscribed", b)
}

func TestEthIsPub(t *testing.T) {
	logging.SetLogLevel("eth-driver", "DEBUG")

	d, err := eth.NewEthDriver()
	if err != nil {
		t.Error(err)
	}

	id, err := peer.IDB58Decode("QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9")
	if err != nil {
		t.Error(err)
	}

	b, err := d.IsPeerAPublisher(id, "t")
	if err != nil {
		t.Error(err)
	}

	t.Log("IsPeerAPublisher", b)
}
