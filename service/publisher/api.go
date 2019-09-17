package publisher

import (
	"errors"
)

// RegisterToPublish helps register to multiple topics at once.
func (pubService *PublisherService) RegisterToPublish(topics []string) error {

	// Temp disabled RA service
	if pubService.ra == nil {
		return nil
	}

	host := pubService.GetHost()

	for _, topic := range topics {
		// TODO: This is a blocking call and won't return until trasaction is finalised.
		// Better to parallalise it somehow.
		err := pubService.ra.Publish(host.ID(), topic)
		if err != nil {
			return err
		}
	}

	return nil
}

// PublishData takes some data and publishes it to all given topics.
func (pubService *PublisherService) PublishData(data []byte, topics []string) error {
	pubRouter := pubService.GetPubSub()
	if pubRouter == nil {
		return errors.New("Publisher is null")
	}

	for _, topic := range topics {
		err := pubRouter.Publish(topic, data)
		if err != nil {
			return err
		}
		log.Info("Pub Size:", len(data))
	}

	return nil
}
