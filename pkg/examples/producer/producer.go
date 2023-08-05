package producer

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func Send() {
	// Set the configuration for the Kafka producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatal("Error creating Kafka producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal("Error closing Kafka producer:", err)
		}
	}()

	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {

		// Create a new Kafka message
		message := &sarama.ProducerMessage{
			Timestamp: time.Now(),
			Topic:     "test",
			Value:     sarama.StringEncoder("Hello, Kafka!"),
		}

		// Send the message to the Kafka broker
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatal("Error sending Kafka message:", err)
		}

		log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	}
}
