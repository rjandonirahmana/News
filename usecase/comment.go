package usecase

import (
	"context"

	"github.com/google/uuid"
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
	id := uuid.New().String()
	comment.ID = id
	err := u.repo.CraeteComment(newsID, comment, ctx)
	if err != nil {
		return err
	}

	return nil
}
