package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"hotel-payments/internal/models"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

var (
	brokers  = os.Getenv("KAFKA_HOSTS")
	version  = sarama.DefaultVersion.String()
	group    = "BOOKING:PAYMENTS"
	topics   = os.Getenv("KAFKA_TOPIC_INITIATE_BOOKING")
	assignor = "roundrobin"
	oldest   = false
	verbose  = false
)

func ServeKafkaConsumer() {
	keepRunning := true
	log.Println("Starting a new Sarama consumer")

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	//consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
			//case <-sigusr1:
			//	toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	d := DependencyInjection()
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			var newBooking models.Booking
			err := json.Unmarshal(message.Value, &newBooking)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				return err
			}

			err = d.PaymentsAPI.ProcessPayment(&newBooking)
			if err != nil {
				log.Printf("failed to process payment: %v", err)
				return err
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
