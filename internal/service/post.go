package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/acool-kaz/parser-service-server/internal/models"
	"github.com/acool-kaz/parser-service-server/internal/repository"
	"golang.org/x/sync/errgroup"
)

type PostService struct {
	webClient *http.Client
	postRepo  repository.Post
}

func newPostService(postRepo repository.Post) *PostService {
	return &PostService{
		webClient: http.DefaultClient,
		postRepo:  postRepo,
	}
}

func (s *PostService) Parse(ctx context.Context, url string, totalPages int) error {
	errs, ctx := errgroup.WithContext(ctx)
	var mu sync.Mutex
	for i := 1; i <= totalPages; i++ {
		func(i int) {
			errs.Go(func() error {
				resp, err := s.webClient.Get(fmt.Sprintf("%s%d", url, i))
				if err != nil {
					return fmt.Errorf("parser service: parse: %w", err)
				}
				defer resp.Body.Close()
				var info models.Info
				if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
					return fmt.Errorf("parser service: parse: %w", err)
				}
				mu.Lock()
				for _, post := range info.Data {
					if err := s.postRepo.Insert(ctx, post); err != nil {
						return fmt.Errorf("parser service: parse: %w", err)
					}
				}
				mu.Unlock()
				return nil
			})
		}(i)
	}

	return errs.Wait()
}
