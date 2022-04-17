package repository

import (
	"context"
	"fmt"

	"github.com/rjandonirahmana/news/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoUser struct {
	db *mongo.Database
}

func NewRepoUser(db *mongo.Database) *repoUser {
	return &repoUser{db: db}
}

type RepoUser interface {
	CreateUser(user *models.User, ctx context.Context) error
	IsEmailAvailable(email *string, ctx context.Context) error
	GetUserByEmail(email *string, ctx context.Context) (*models.User, error)
	GetUserByID(id *string, ctx context.Context) (*models.User, error)
}

func (r *repoUser) CreateUser(user *models.User, ctx context.Context) error {
	if _, err := r.db.Collection("users").InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (r *repoUser) IsEmailAvailable(email *string, ctx context.Context) error {
	rslt := r.db.Collection("users").FindOne(ctx, bson.M{"email": *email})
	var user models.User
	err := rslt.Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err != nil {
		return err
	}

	return fmt.Errorf("email has been used")
}

func (r *repoUser) GetUserByEmail(email *string, ctx context.Context) (*models.User, error) {
	rslt := r.db.Collection("users").FindOne(ctx, bson.M{"email": *email})
	var user *models.User

	err := rslt.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repoUser) GetUserByID(id *string, ctx context.Context) (*models.User, error) {
	rslt := r.db.Collection("users").FindOne(ctx, bson.M{"id": *id})
	var user *models.User

	err := rslt.Decode(&user)

	if err == mongo.ErrNoDocuments {
		return user, fmt.Errorf("user not found")
	}

	if err != nil {
		return user, err
	}

	return user, nil
}
