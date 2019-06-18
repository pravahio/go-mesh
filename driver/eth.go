package driver

import peer "github.com/libp2p/go-libp2p-core/peer"

type EthDriver struct{}

func NewEthDriver() *EthDriver {
	return &EthDriver{}
}

func (eth *EthDriver) Subscribe() error {
	return nil
}

func (eth *EthDriver) Publish() error {
	return nil
}

func (eth *EthDriver) IsPeerAPublisher(p peer.ID, t string) bool {
	/* if p.Pretty() == "QmRq6L8f7rPpKPNJXxovN4ZhUsNgWzXWDmvzBTS8hJQ5C1" {
		return false
	} */
	return true
}

func (eth *EthDriver) IsPeerSubscribed(p peer.ID, t string) bool {
	/* if p.Pretty() == "QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9" {
		return false
	} */
	return true
}
