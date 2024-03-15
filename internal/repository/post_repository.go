package repository

import (
	"context"

	"github.com/you-5805/learning-go/internal/domain"
)

type PostRepository interface {
	Fetch(ctx context.Context) ([]domain.Post, error)
	Store(ctx context.Context, post *domain.Post) error
}
