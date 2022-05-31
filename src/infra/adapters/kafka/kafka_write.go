package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	write *kafka.Writer
}

func NewKafkaClient(write *kafka.Writer) KafkaClient {
	return KafkaClient{
		write: write,
	}
}

func (k KafkaClient) PublishMessage(ctx context.Context, id string, message interface{}, topic string, headers map[string]string) error {
	messageEncoded, err := json.Marshal(message)
	if err != nil {
		return errors.New("Error on publish message: " + err.Error())
	}

	var kafkaHeaders []kafka.Header

	err = k.write.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:     []byte(id),
			Value:   messageEncoded,
			Headers: kafkaHeaders,
			Topic:   topic,
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("k.writer.writeMessages: " + err.Error())
	}
	return nil
}
