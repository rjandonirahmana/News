package repository

import (
	"context"

	"github.com/rjandonirahmana/news/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoComment struct {
	db *mongo.Database
}

func NewRepoComments(db *mongo.Database) *repoComment {
	return &repoComment{db: db}
}

type RepoComment interface {
	UpdateCommentComment(newsID *string, comment *models.CommentNews, ctx context.Context) error
	CraeteComment(newsID *string, comment *models.CommentNews, ctx context.Context) error
}

//update news in comment value
func (r *repoComment) UpdateCommentComment(newsID *string, comment *models.CommentNews, ctx context.Context) error {
	filter := bson.M{"id": *newsID, "comments": bson.M{"$elemMatch": bson.M{"user_id": bson.M{"$lte": comment.UserID}}}}
	// update := bson.M{"$set" : }
	_, err := r.db.Collection("news").UpdateOne(ctx, filter, bson.M{"$set": comment})

	if err != nil {
		return err
	}

	return nil
}

func (r *repoComment) CraeteComment(newsID *string, comment *models.CommentNews, ctx context.Context) error {
	filter := bson.M{"id": *newsID}
	// update := bson.M{"$set" : }
	_, err := r.db.Collection("news").UpdateOne(ctx, filter, bson.M{"$set": bson.M{"comments": comment}})

	if err != nil {
		return err
	}

	return nil
}
