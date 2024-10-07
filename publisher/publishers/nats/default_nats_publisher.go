package nats

import (
	"encoding/json"
	"fmt"
	pub "publish-expcetion/publisher/publishers"

	"github.com/nats-io/nats.go"
)

type DefaultNatsPublisher struct{}

func (DefaultNatsPublisher) Publish(msg *pub.MessageException) error {
	// Connect to the NATS server
	conn, err := nats.Connect("localhost")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Encode the message object as JSON
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Define the NATS topic to publish to
	topic := msg.ApplicationName + "-publish-exception"

	// Create a JetStream context
	js, err := conn.JetStream()
	if err != nil {
		return fmt.Errorf("failed to create JetStream context: %w", err)
	}

	// Publish the message to the NATS server
	_, err = js.PublishAsync(topic, data)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	fmt.Printf("Published message to NATS topic: %s message:%s\n", topic, data)

	return nil
}
