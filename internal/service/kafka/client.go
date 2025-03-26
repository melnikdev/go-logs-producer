package kafka

import (
	"context"
	"encoding/json"

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
		Addr:     kafka.TCP(c.config.GRPCServer.Addr),
		Topic:    c.config.GRPCServer.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	messageBytes, _ := json.Marshal(logMessage)
	return writer.WriteMessages(context.Background(), kafka.Message{Value: messageBytes})
}
