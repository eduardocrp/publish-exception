package kafka

import (
	"fmt"
	pub "publish-expcetion/publisher/publisher"
)

type DefaultKafkaPublisher struct{}

func (DefaultKafkaPublisher) Publish(msg *pub.MessageException) error {
	fmt.Printf("Teste default kafka publisher: %v\n", msg)
	return nil
}
