package subscriber

import (
	"errors"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// SubscribeToTopics subscribes to multiple topics and outputs all into a single chan.
func (subService *SubscriberService) SubscribeToTopics(topics []string) (chan *pubsub.Message, error) {
	// Temp disabled RA service
	if subService.ra != nil {

		// 1. Register on the blockchain
		// 2. Subscribe to pubsub

		host := subService.GetHost()

		for _, topic := range topics {
			err := subService.ra.Subscribe(host.ID(), string(topic))
			if err != nil {
				return nil, err
			}
		}
	}

	// Transaction on the blockchain won't be reflect immediately.
	// TODO: So, wait before sending a sub message.

	subRouter := subService.GetPubSub()
	if subRouter == nil {
		return nil, errors.New("subscriber router is nil")
	}

	msg := make(chan *pubsub.Message, 100)

	for _, topic := range topics {

		sub, err := subRouter.Subscribe(topic)
		if err != nil {
			return nil, err
		}

		subService.topicTracker[topic] = &TopicWrapper{
			subscription: sub,
		}

		go func(subs *pubsub.Subscription) {
			for {
				select {
				case <-subService.ctx.Done():
					close(msg)
					return
				default:
				}

				m, err := subs.Next(subService.ctx)
				if err != nil {
					log.Error(err)
					close(msg)
					return
				}

				msg <- m
			}

		}(sub)
	}

	return msg, nil
}

// UnsubscribeToTopics will unregister from the remote access and close the subscription channel
func (subService *SubscriberService) UnsubscribeToTopics(topics []string) error {
	if subService.ra != nil {

		// 1. Register on the blockchain
		// 2. Subscribe to pubsub

		host := subService.GetHost()

		for _, topic := range topics {
			err := subService.ra.Unsubscribe(host.ID(), string(topic))
			if err != nil {
				return err
			}
		}
	}

	for _, topic := range topics {
		tw, ok := subService.topicTracker[topic]
		if !ok {
			return errors.New("topic not found in the list of subscribed topics")
		}
		tw.subscription.Cancel()
		subService.topicTracker[topic] = nil
	}

	return nil
}
