package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/repository"
)

type usecaseNews struct {
	repo repository.RepoNews
}

type UsecaseNews interface {
	CreateNews(news *models.News, ctx context.Context) (*models.News, error)
	GetNewsByID(id *string, ctx context.Context) (*models.News, error)
	DeleteNewsByID(newsID *string, ctx context.Context) error
}

func NewServiceNews(repo repository.RepoNews) *usecaseNews {
	return &usecaseNews{repo: repo}
}

func (u *usecaseNews) CreateNews(news *models.News, ctx context.Context) (*models.News, error) {
	news.ID = uuid.New().String()
	news.Images = []models.Images{}
	news.Comment = []models.CommentNews{}
	news.CreatedAt = time.Now()
	news.UpdatedAt = time.Now()
	news, err := u.repo.CreateNews(news, ctx)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (u *usecaseNews) GetNewsByID(id *string, ctx context.Context) (*models.News, error) {
	return u.repo.GetNewsByID(id, ctx)
}

func (u *usecaseNews) DeleteNewsByID(newsID *string, ctx context.Context) error {
	return u.repo.DeleteNewsByID(newsID, ctx)
}
