package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server     *Server
	GRPCServer *GRPCServer
}

type Server struct {
	Port    string
	Network string
}

type GRPCServer struct {
	Addr  string
	Topic string
}

func NewConfig() *Config {
	return &Config{
		Server: &Server{
			Network: getEnv("SERVER_NETWORK", "tcp"),
			Port:    getEnv("SERVER_PORT", "8080"),
		},
		GRPCServer: &GRPCServer{
			Addr:  getEnv("GRPC_SERVER_ADDR", "localhost:9092"),
			Topic: getEnv("GRPC_SERVER_TOPIC", "logs_topic"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
