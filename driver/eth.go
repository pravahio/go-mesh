package driver

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
