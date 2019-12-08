package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	types "github.com/ethereum/go-ethereum/core/types"
	crypto "github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	ra "github.com/pravahio/go-mesh/interface/ra"
)

var log = logging.Logger("eth-driver")

const TRIES = 10

var triesInterval = 5 * time.Second

// EthDriver implements "RemoteAccess".
type EthDriver struct {
	client *ethclient.Client
	ether  *Eth
	key    *ecdsa.PrivateKey
	opt    *bind.TransactOpts
}

func NewEthDriver(URI string, prvKey *ecdsa.PrivateKey) (ra.RemoteAccess, error) {

	if prvKey == nil {
		var err error
		prvKey, err = CreateNewAccount()
		if err != nil {
			return nil, err
		}
		fmt.Println("Creating a new account. It might not be able to do any transaction.")
		log.Info("Creating a new account.")
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
		return err
	}
	log.Info("Trans Hash: ", t.Hash().String())

	for i := 0; i < TRIES; i++ {
		rec, err := eth.client.TransactionReceipt(context.Background(), t.Hash())
		if err == nil {
			log.Info("Transaction is success", rec.Status == types.ReceiptStatusSuccessful)
			return nil
		}
		log.Info("Transaction pending", err)

		time.Sleep(triesInterval)
	}

	return errors.New("Something happened with the transaction")
}

func (eth *EthDriver) Unsubscribe(p peer.ID, topic string) error {

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

	t, err := eth.ether.Unsubscribe(eth.opt, a, topic)
	if err != nil {
		return err
	}
	log.Info("Trans Hash: ", t.Hash().String())

	for i := 0; i < TRIES; i++ {
		rec, err := eth.client.TransactionReceipt(context.Background(), t.Hash())
		if err == nil {
			log.Info("Transaction is success", rec.Status == types.ReceiptStatusSuccessful)
			return nil
		}
		log.Info("Transaction pending", err)

		time.Sleep(triesInterval)
	}

	return errors.New("Something happened with the transaction")
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

	for i := 0; i < TRIES; i++ {
		rec, err := eth.client.TransactionReceipt(context.Background(), t.Hash())
		if err == nil {
			log.Info("Transaction is success", rec.Status == types.ReceiptStatusSuccessful)
			return nil
		}
		log.Info("Transaction pending", err)

		time.Sleep(triesInterval)
	}

	return errors.New("Something happened with the transaction")
}

func (eth *EthDriver) Unpublish(p peer.ID, topic string) error {
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

	t, err := eth.ether.Unpublish(eth.opt, a, topic)
	if err != nil {
		return err
	}
	log.Info("Trans Hash: ", t.Hash().String())

	for i := 0; i < TRIES; i++ {
		rec, err := eth.client.TransactionReceipt(context.Background(), t.Hash())
		if err == nil {
			log.Info("Transaction is success", rec.Status == types.ReceiptStatusSuccessful)
			return nil
		}
		log.Info("Transaction pending", err)

		time.Sleep(triesInterval)
	}

	return errors.New("Something happened with the transaction")
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
	auth.GasLimit = uint64(600000)

	return auth
}
