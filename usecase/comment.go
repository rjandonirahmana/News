package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/repository"
)

type usecaseComment struct {
	repo repository.RepoComment
}

type UsecaseComment interface {
	CreateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error
	LikeComment(newsID *string, commentID *string, ctx context.Context) error
}

func NewUseCaseComent(repo repository.RepoComment) *usecaseComment {
	return &usecaseComment{repo: repo}
}

func (u *usecaseComment) CreateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error {
	id := uuid.New().String()
	comment.ID = id
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	err := u.repo.CreateComment(newsID, comment, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecaseComment) LikeComment(newsID *string, commentID *string, ctx context.Context) error {
	return u.repo.AddLikeComment(newsID, commentID, ctx)
}
