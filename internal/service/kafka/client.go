package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/melnikdev/go-logs-producer/internal/config"
	"github.com/segmentio/kafka-go"
)

type LogMessage struct {
	Service   string `json:"service"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type ClientI interface {
	Send(logMessage LogMessage) error
}

type client struct {
	config *config.Config
}

func NewKafkaClient(config *config.Config) ClientI {
	return &client{config: config}
}

func (c *client) Send(logMessage LogMessage) error {

	writer := &kafka.Writer{
		Addr:     kafka.TCP(c.config.KAFKA.Broker),
		Topic:    c.config.KAFKA.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	messageBytes, _ := json.Marshal(logMessage)
	fmt.Println(string(messageBytes))
	return writer.WriteMessages(context.Background(), kafka.Message{Value: messageBytes})
}
