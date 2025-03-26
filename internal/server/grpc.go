package server

import (
	"context"

	"github.com/melnikdev/go-logs-producer/internal/service/kafka"
	pb "github.com/melnikdev/go-logs-producer/proto"
)

type server struct {
	kafkaClient kafka.ClientI
	pb.UnimplementedLogServiceServer
}

func NewLogGRPCServer(kafkaClient kafka.ClientI) pb.LogServiceServer {
	return &server{kafkaClient: kafkaClient}
}

func (s *server) SendLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {

	logMessage := kafka.LogMessage{
		Service:   req.Service,
		Level:     req.Level,
		Message:   req.Message,
		Timestamp: req.Timestamp,
	}

	err := s.kafkaClient.Send(logMessage)

	if err != nil {
		return &pb.LogResponse{Status: "error"}, err
	}

	return &pb.LogResponse{Status: "success"}, nil
}
