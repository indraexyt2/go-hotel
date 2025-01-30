package cmd

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"hotel-payments/helpers"
	"log"
	"os"
	"strings"
	"sync"
)

type Consumer struct {
	Ready chan bool
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.Ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			helpers.Logger.Info(fmt.Sprintf("Message topic:%s partition:%d offset:%d", message.Topic, message.Partition, message.Offset))

			fmt.Printf("Message value: %s\n", string(message.Value))

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func ServeKafkaConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := strings.Split(os.Getenv("KAFKA_HOSTS"), ",")
	topics := os.Getenv("KAFKA_TOPIC_INITIATE_BOOKING")
	group := "BOOKING:PAYMENTS"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumer := Consumer{
		Ready: make(chan bool),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if ctx.Err() != nil {
				return
			}
			err := client.Consume(ctx, []string{topics}, &consumer)
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
			consumer.Ready = make(chan bool)
		}
	}()

	<-consumer.Ready

	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}
