package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/melnikdev/go-logs-producer/internal/config"
	"github.com/melnikdev/go-logs-producer/internal/server"
	"github.com/melnikdev/go-logs-producer/internal/service/kafka"
	pb "github.com/melnikdev/go-logs-producer/proto"

	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.NewConfig()

	listener, err := net.Listen(config.Server.Network, ":"+config.Server.Port)

	if err != nil {
		log.Fatalf("Error run server: %v", err)
	}

	kafkkClient := kafka.NewKafkaClient(config)
	logServer := server.NewLogGRPCServer(kafkkClient)

	grpcServer := grpc.NewServer()
	pb.RegisterLogServiceServer(grpcServer, logServer)

	fmt.Printf("gRPC listener %v...", config.Server.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error run grpc server: %v", err)
	}
}
