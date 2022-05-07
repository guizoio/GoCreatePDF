package kafka

import (
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

func NewKafkaWriter(broker string, useTLS bool) *kafka.Writer {
	return kafka.NewWriter(
		kafka.WriterConfig{
			Brokers: strings.Split(broker, ","),
			Dialer:  getDialer(useTLS),
		})
}

func getDialer(userTLS bool) *kafka.Dialer {
	if !userTLS {
		return &kafka.Dialer{}
	}

	return &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{},
	}
}
