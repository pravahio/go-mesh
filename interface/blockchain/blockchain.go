package blockchain

type Blockchain interface {
	// Subscriber specific
	Subscribe() error

	// Publisher specific
	Publish() error
}