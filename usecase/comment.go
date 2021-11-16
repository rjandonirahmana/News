package usecase

import (
	"context"

	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/repository"
)

type usecaseComment struct {
	repo repository.RepoComment
}

type UsecaseComment interface {
	CreateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error
}

func NewUseCaseComent(repo repository.RepoComment) *usecaseComment {
	return &usecaseComment{repo: repo}
}

func (u *usecaseComment) CreateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error {
	return u.repo.CraeteComment(newsID, comment, ctx)
}
