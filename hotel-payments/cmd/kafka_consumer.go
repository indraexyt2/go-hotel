package cmd

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"hotel-payments/helpers"
	"hotel-payments/internal/models"
	"os"
	"strconv"
	"strings"
)

func ServeKafkaConsumer() {
	d := DependencyInjection()
	brokers := strings.Split(os.Getenv("KAFKA_HOSTS"), ",")
	topic := os.Getenv("KAFKA_TOPIC_INITIATE_BOOKING")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Error creating Kafka consumer: ", err)
		return
	}

	partitionNumberStr := os.Getenv("KAFKA_PARTITION_NUMBER")
	partitionNumber, _ := strconv.Atoi(partitionNumberStr)
	for i := 0; i < partitionNumber; i++ {
		go func() {
			partitionConsumer, err := consumer.ConsumePartition(topic, int32(i), sarama.OffsetNewest)
			if err != nil {
				helpers.Logger.Error("Error consuming Kafka partition: ", err)
				return
			}

			for msg := range partitionConsumer.Messages() {
				helpers.Logger.Info("Received message payment processed by consumer: ", string(msg.Value))

				var reqBooking models.Booking
				err = json.Unmarshal(msg.Value, &reqBooking)
				if err != nil {
					helpers.Logger.Error("Error unmarshalling Kafka message: ", err)
					return
				}

				err = d.PaymentsAPI.ProcessPayment(&reqBooking)
				if err != nil {
					helpers.Logger.Error("Error processing payment: ", err)
					return
				}
			}
		}()
	}
}
