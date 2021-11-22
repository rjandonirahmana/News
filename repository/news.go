package repository

import (
	"context"
	"fmt"
	"log"

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
		fmt.Println(err)
		return nil, err

	}

	return news, nil

}

func (r *repoNews) GetNewsByID(id *string, ctx context.Context) (*models.News, error) {
	rslt := r.db.Collection("news").FindOne(ctx, bson.M{"_id": *id})
	var news *models.News
	err := rslt.Decode(&news)

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r *repoNews) UpdateNews(news *models.News, ctx context.Context) (*models.News, error) {
	filter := bson.M{"_id": news.ID}
	update := bson.M{
		"$set": bson.M{
			"title":      news.Name,
			"author":     news.Author,
			"updated_at": news.UpdatedAt,
			"content":    news.Content,
		},
	}

	_, err := r.db.Collection("news").UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return news, nil

}

func (r *repoNews) DeleteALLNews(ctx context.Context) error {

	filter := bson.M{}
	_, err := r.db.Collection("news").DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")

	return nil
}

func (r *repoNews) UpdatePhotoNews(file string, id string, ctx context.Context) (*models.News, error) {
	return &models.News{}, nil
}
