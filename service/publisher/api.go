package publisher

import (
	"errors"
)

func (pubService *PublisherService) RegisterToPublish(topic string) error {
	host := pubService.GetHost()
	err := pubService.ra.Publish(host.ID(), topic)
	if err != nil {
		return err
	}

	return nil
}

func (pubService *PublisherService) PublishData(topic string, data []byte) error {
	pubRouter := pubService.GetPubSub()
	if pubRouter == nil {
		return errors.New("Publisher is null")
	}
	err := pubRouter.Publish(topic, data)
	if err != nil {
		return err
	}
	log.Info("Pub", string(data))

	return nil
}
