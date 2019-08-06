package subscriber

import (
	"errors"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func (subService *SubscriberService) SubscribeToTopic(topic string) (chan *pubsub.Message, error) {
	// 1. Register on the blockchain
	// 2. Subscribe to pubsub

	host := subService.GetHost()
	err := subService.ra.Subscribe(host.ID(), string(topic))
	if err != nil {
		return nil, err
	}

	// Transaction on the blockchain won't be reflect immediately.
	// TODO: So, wait before sending a sub message.

	subRouter := subService.GetPubSub()
	if subRouter == nil {
		return nil, errors.New("subscriber router is nil")
	}

	sub, err := subRouter.Subscribe(topic)
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
