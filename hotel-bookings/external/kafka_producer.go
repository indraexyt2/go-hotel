package external

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"hotel-bookings/helpers"
	"os"
	"strings"
	"time"
)

func (ex *External) ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error {
	brokers := strings.Split(os.Getenv("KAFKA_HOSTS"), ",")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Timeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return err
	}

	helpers.Logger.Info(fmt.Sprintf("message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset))
	return nil
}
