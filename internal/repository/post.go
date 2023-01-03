package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/acool-kaz/parser-service-server/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func newPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Insert(ctx context.Context, post models.Post) error {
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, title, body) VALUES($1, $2, $3, $4);", postTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("post repository: insert: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, post.Id, post.UserId, post.Title, post.Body); err != nil {
		return fmt.Errorf("post repository: insert: %w", err)
	}

	return nil
}
