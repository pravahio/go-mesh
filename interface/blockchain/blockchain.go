package blockchain

type Blockchain interface {
	// Subscriber specific
	Subscribe()

	// Publisher specific
	Publish()
}