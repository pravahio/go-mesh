package publisher

func (pubService *PublisherService) RegisterToPublish(topic TopicName) error {
	err := pubService.blockchain.Publish()
	if err != nil {
		return err
	}

	return nil
}

func (pubService *PublisherService) PublishData(topic TopicName, data []byte) error {
	err := pubService.GetPubSub().Publish(string(topic), data)
	if err != nil {
		return err
	}
	log.Info("Pub", string(data))

	return nil
}
