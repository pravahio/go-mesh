package eth

import (
	"context"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

var (
	contractStaticAddress = "0xf0d731661b7cd29a6677f8255d74e075f412a77d"
	ethClientAddress      = "http://127.0.0.1:7545"
)

func loadContract(URI string) (*ethclient.Client, *Eth, error) {
	if URI != "" {
		ethClientAddress = URI
	}
	c, err := ethclient.Dial(ethClientAddress)
	if err != nil {
		return nil, nil, err
	}

	netID, err := c.NetworkID(context.Background())
	if err != nil {
		return nil, nil, err
	}
	log.Info(netID)

	contractAddress := common.HexToAddress(contractStaticAddress)

	eth, err := NewEth(contractAddress, c)
	if err != nil {
		return nil, nil, err
	}

	return c, eth, nil
}

func libp2pToEthAccount(p peer.ID) (common.Address, error) {
	bKey, err := p.MarshalBinary()
	if err != nil {
		return common.Address{}, err
	}
	a := common.BytesToAddress(bKey)

	return a, nil
}
