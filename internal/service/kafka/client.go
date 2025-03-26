package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type LogMessage struct {
	Service   string `json:"service"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type Client interface {
	Send(logMessage LogMessage) error
}

type client struct {
}

func NewKafkaClient() Client {
	return &client{}
}

func (c *client) Send(logMessage LogMessage) error {

	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "logs_topic",
		Balancer: &kafka.LeastBytes{},
	}

	messageBytes, _ := json.Marshal(logMessage)
	return writer.WriteMessages(context.Background(), kafka.Message{Value: messageBytes})
}
