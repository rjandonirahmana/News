package repository

import (
	"context"

	"github.com/rjandonirahmana/news/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoUser struct {
	db *mongo.Database
}

func NewRepoUser(db *mongo.Database) *repoUser {
	return &repoUser{db: db}
}

func (r *repoUser) CreateUser(user *models.User, ctx context.Context) error {
	if _, err := r.db.Collection("users").InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

// func (r *repoUser) GetByID
