package main

import (
	"fmt"
	"log"
	"net"

	"github.com/melnikdev/go-logs-producer/internal/server"
	pb "github.com/melnikdev/go-logs-producer/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Error run server: %v", err)
	}

	logServer := server.NewLogGRPCServer()
	grpcServer := grpc.NewServer()
	pb.RegisterLogServiceServer(grpcServer, logServer)

	fmt.Println("gRPC listener 50051...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error run grpc server: %v", err)
	}
}
