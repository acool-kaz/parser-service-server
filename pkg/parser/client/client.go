package client

import (
	"fmt"
	"log"

	parser_pb "github.com/acool-kaz/parser-service-server/pkg/parser/pb"
	"google.golang.org/grpc"
)

type ParserClientConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func InitParserClientConfig(host, port string) *ParserClientConfig {
	log.Println("init parser client config")

	return &ParserClientConfig{
		Host: host,
		Port: port,
	}
}

type ParserClient struct {
	client parser_pb.ParserServiceClient
}

func InitParserClient(cfg *ParserClientConfig) (*ParserClient, error) {
	log.Println("init parser client")

	conn, err := grpc.Dial(cfg.Host+":"+cfg.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("init parser client: %w", err)
	}

	return &ParserClient{
		client: parser_pb.NewParserServiceClient(conn),
	}, nil
}
