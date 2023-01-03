package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/acool-kaz/parser-service-server/internal/models"
)

const postTable = "posts"

type Post interface {
	Insert(ctx context.Context, post models.Post) error
}

type Repository struct {
	Post Post
}

func InitRepository(db *sql.DB) *Repository {
	log.Println("init repository")
	return &Repository{
		Post: newPostRepository(db),
	}
}
