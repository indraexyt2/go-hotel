package external

import (
	"context"
	"github.com/IBM/sarama"
	"hotel-ums/helpers"
	"os"
	"strings"
	"time"
)

type External struct{}

func NewExternal() *External {
	return &External{}
}

func (ex *External) SendMessageNotification(ctx context.Context, message []byte) error {
	brokers := strings.Split(os.Getenv("KAFKA_HOSTS"), ",")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Timeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Failed to create Kafka producer: ", err)
		return err
	}

	kafkaMessage := &sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC_NOTIFICATION"),
		Value: sarama.ByteEncoder(message),
	}

	_, _, err = producer.SendMessage(kafkaMessage)
	if err != nil {
		helpers.Logger.Error("Failed to send message to Kafka: ", err)
		return err
	}

	return nil
}
