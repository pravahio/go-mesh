package bcinterface

type BlockchainDriver interface {
	// Subscriber specific
	Subscribe()

	// Publisher specific
	Publish()
}