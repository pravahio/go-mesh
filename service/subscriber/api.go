package subscriber

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func (subService *SubscriberService) SubscribeToTopic(topic TopicName) (chan *pubsub.Message, error) {
	// 1. Register on the blockchain
	// 2. Subscribe to pubsub

	err := subService.ra.Subscribe()
	if err != nil {
		return nil, err
	}

	// Transaction on the blockchain won't be reflect immediately.
	// TODO: So, wait before sending a sub message.

	sub, err := subService.GetPubSub().Subscribe(string(topic))
	if err != nil {
		return nil, err
	}

	subService.topicTracker[topic] = &TopicWrapper{
		subscription: sub,
	}

	msg := make(chan *pubsub.Message, 100)

	go func() {
		for {
			select {
			case <-subService.ctx.Done():
				close(msg)
				return
			default:
			}

			m, err := sub.Next(subService.ctx)
			if err != nil {
				continue
			}

			msg <- m
		}

	}()

	return msg, nil
}
