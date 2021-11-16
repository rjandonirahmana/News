package repository

import (
	"context"

	"github.com/rjandonirahmana/news/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoNews struct {
	db *mongo.Database
}

func NewRepoNews(db *mongo.Database) *repoNews {
	return &repoNews{db: db}
}

type RepoNews interface {
	CreateNews(news *models.News, ctx context.Context) (*models.News, error)
	GetNewsByID(id *string, ctx context.Context) (*models.News, error)
}

func (r *repoNews) CreateNews(news *models.News, ctx context.Context) (*models.News, error) {
	_, err := r.db.Collection("news").InsertOne(ctx, news)

	if err != nil {
		return nil, err
	}

	return news, nil

}

func (r *repoNews) GetNewsByID(id *string, ctx context.Context) (*models.News, error) {
	rslt := r.db.Collection("news").FindOne(ctx, bson.M{"id": *id})
	var news *models.News
	err := rslt.Decode(&news)

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r *repoNews) UpdateNews(news *models.News, ctx context.Context) (*models.News, error) {
	selector := bson.M{"id": news.ID}
	_, err := r.db.Collection("news").UpdateOne(ctx, selector, bson.M{"$set": news})

	if err != nil {
		return nil, err
	}

	return news, nil

}

func (r *repoNews) UpdatePhotoNews(file string, id string, ctx context.Context) (*models.News, error) {
	return &models.News{}, nil
}
