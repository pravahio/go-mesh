package subscriber

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type TopicWrapper struct {
	subscription *pubsub.Subscription
}
