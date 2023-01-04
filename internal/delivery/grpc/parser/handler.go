package parser

import (
	"context"
	"log"

	"github.com/acool-kaz/parser-service-server/internal/config"
	"github.com/acool-kaz/parser-service-server/internal/service"
	parser_pb "github.com/acool-kaz/parser-service-server/pkg/parser/pb"
)

type ParserHandler struct {
	cfg *config.Config
	parser_pb.UnimplementedParserServiceServer
	service *service.Service
}

func InitParserHandler(cfg *config.Config, service *service.Service) *ParserHandler {
	log.Println("init grpc parser handler")
	return &ParserHandler{
		cfg:     cfg,
		service: service,
	}
}

func (p *ParserHandler) Parser(ctx context.Context, req *parser_pb.ParserRequest) (*parser_pb.ParserResponse, error) {
	if err := p.service.Post.Parse(ctx, p.cfg.Parser.ParseUrl, p.cfg.Parser.MaxPage); err != nil {
		return &parser_pb.ParserResponse{Status: "Failed"}, err
	}

	return &parser_pb.ParserResponse{Status: "OK"}, nil
}
