package consumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

func Start(cancel chan string) error {
	// Set the configuration for the Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Printf("Error creating Kafka consumer: %s", err)
		return err
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("Error closing Kafka consumer: %s", err)
		}
	}()

	// Create a new Kafka partition consumer for the "test" topic and partition 0
	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		log.Printf("Error creating Kafka partition consumer: %s", err)
		return err
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Printf("Error creating Kafka partition consumer: %s", err)
		}
	}()

	// Handle Kafka messages in a separate goroutine
	go func() {
		for message := range partitionConsumer.Messages() {
			fmt.Printf("Message received: topic=%s partition=%d offset=%d value=%s\n", message.Topic, message.Partition, message.Offset, string(message.Value))
		}
	}()

	fmt.Println("----------Waiting for cutoff signal")
	// Wait for an interrupt signal (e.g. Ctrl+C) to stop the consumer
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	fmt.Println("Kafka consumer stopped")
	return nil
}

func BackgroundRoutine(ch <-chan string) {
	for {
		// Wait for a message on the channel
		Start(make(chan string))

		// Process the message in the background
		fmt.Println("Consumer died... Restarting in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
