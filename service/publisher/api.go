package publisher

func (pubService *PublisherService) RegisterToPublish(topic string) error {
	host := pubService.GetHost()
	err := pubService.ra.Publish(host.ID(), topic)
	if err != nil {
		return err
	}

	return nil
}

func (pubService *PublisherService) PublishData(topic string, data []byte) error {
	err := pubService.GetPubSub().Publish(topic, data)
	if err != nil {
		return err
	}
	log.Info("Pub", string(data))

	return nil
}
