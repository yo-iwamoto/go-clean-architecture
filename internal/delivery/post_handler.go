package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/you-5805/learning-go/internal/domain"
	"github.com/you-5805/learning-go/internal/usecase"
)

type PostHandler struct {
	PUsecase usecase.PostUsecase
}

func NewPostHandler(pu usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		PUsecase: pu,
	}
}

func (p *PostHandler) GetPosts(c echo.Context) {
	posts, err := p.PUsecase.FetchPosts(c)
	writer := c.Response().Writer
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(posts)
}

func (p *PostHandler) CreatePost(c echo.Context) {
	var post domain.Post
	writer := c.Response().Writer
	err := json.NewDecoder(c.Request().Body).Decode(&post)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.PUsecase.CreatePost(c, &post)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
