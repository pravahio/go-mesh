package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	crypto "github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	ra "github.com/upperwal/go-mesh/interface/ra"
)

var log = logging.Logger("eth-driver")

// EthDriver implements "RemoteAccess".
type EthDriver struct {
	client *ethclient.Client
	ether  *Eth
	key    *ecdsa.PrivateKey
	opt    *bind.TransactOpts
}

func NewEthDriver(URI string, prvKey *ecdsa.PrivateKey) (ra.RemoteAccess, error) {

	if prvKey == nil {
		return nil, errors.New("Private Key is nil")
	}

	c, e, err := loadContract(URI)
	if err != nil {
		return nil, err
	}

	o := createTransOpt(prvKey)

	return &EthDriver{
		client: c,
		ether:  e,
		key:    prvKey,
		opt:    o,
	}, nil
}

func (eth *EthDriver) Subscribe(p peer.ID, topic string) error {

	fromAddress := crypto.PubkeyToAddress(eth.key.PublicKey)
	log.Info("Pub Address: ", fromAddress.String())
	nonce, err := eth.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	eth.opt.Nonce = big.NewInt(int64(nonce))

	a, err := libp2pToEthAccount(p)
	if err != nil {
		return err
	}
	log.Info("Address: ", a.String())

	t, err := eth.ether.Subscribe(eth.opt, a, topic)
	if err != nil {
		log.Info(t, err)
		return err
	}
	log.Info("Trans Hash: ", t.Hash().String())
	return nil
}

func (eth *EthDriver) Publish(p peer.ID, topic string) error {
	fromAddress := crypto.PubkeyToAddress(eth.key.PublicKey)
	log.Info("Pub Address: ", fromAddress.String())
	nonce, err := eth.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	eth.opt.Nonce = big.NewInt(int64(nonce))

	a, err := libp2pToEthAccount(p)
	if err != nil {
		return err
	}
	log.Info("Address: ", a.String())

	t, err := eth.ether.Publish(eth.opt, a, topic)
	if err != nil {
		return err
	}
	log.Info("Trans Hash: ", t.Hash().String())
	return nil
}

func (eth *EthDriver) IsPeerAPublisher(p peer.ID, t string) (bool, error) {
	/* if p.Pretty() == "QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9" {
		return false, nil
	} */
	a, err := libp2pToEthAccount(p)
	if err != nil {
		return false, err
	}

	v, err := eth.ether.IsPeerAPublisher(nil, a, t)
	if err != nil {
		return false, err
	}
	log.Info("Key: ", a.String(), v)
	return v, nil
}

func (eth *EthDriver) IsPeerSubscribed(p peer.ID, t string) (bool, error) {
	/* if p.Pretty() == "QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9" {
		return false, nil
	} */
	a, err := libp2pToEthAccount(p)
	if err != nil {
		return false, err
	}
	log.Info("Key: ", a.String())

	v, err := eth.ether.IsPeerSubscribed(nil, a, t)
	if err != nil {
		return false, err
	}
	return v, nil
}

func createTransOpt(k *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(k)
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(20000000000)
	auth.GasLimit = uint64(6721975)

	return auth
}
