package usecase

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/you-5805/learning-go/internal/domain"
)

type PostUsecase interface {
	FetchPosts(ctx echo.Context) ([]domain.Post, error)
	CreatePost(ctx echo.Context, post *domain.Post) error
}

type postUsecase struct {
	mu    sync.Mutex
	posts []domain.Post
}

func NewPostUsecase() PostUsecase {
	return &postUsecase{
		posts: []domain.Post{
			{ID: 1, Title: "初投稿", Content: "こんにちは、世界！"},
		},
	}
}

func (p *postUsecase) FetchPosts(ctx echo.Context) ([]domain.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.posts, nil
}

func (p *postUsecase) CreatePost(ctx echo.Context, post *domain.Post) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	post.ID = len(p.posts) + 1
	p.posts = append(p.posts, *post)

	return nil
}
