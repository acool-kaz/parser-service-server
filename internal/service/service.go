package service

import (
	"context"
	"log"

	"github.com/acool-kaz/parser-service-server/internal/repository"
)

type Post interface {
	Parse(ctx context.Context, url string, totalPages int) error
}

type Service struct {
	Post Post
}

func InitService(repo *repository.Repository) *Service {
	log.Println("init service")
	return &Service{
		Post: newPostService(repo.Post),
	}
}
