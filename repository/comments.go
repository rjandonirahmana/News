package repository

import (
	"context"
	"fmt"

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
	UpdateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error
	CreateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error
	DeleteComment(newsID *string, commentID *string, ctx context.Context) error
	AddLikeComment(newsID, commentID *string, ctx context.Context) error
}

//update news in comment value
func (r *repoComment) UpdateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error {
	filter := bson.M{"_id": *newsID, "comments._id": comment.ID}
	update := bson.M{
		"$set": bson.M{
			"comments.$.comment":    comment.Comment,
			"comments.$.updated_at": comment.UpdatedAt,
		},
	}
	rslt, err := r.db.Collection("news").UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if rslt.MatchedCount == 0 {
		fmt.Println("ga update")
	}

	return nil
}

func (r *repoComment) CreateComment(newsID *string, comment *models.CommentNews, ctx context.Context) error {
	filter := bson.M{"_id": *newsID}
	update := bson.M{
		"$push": bson.M{
			"comments": bson.M{
				"_id":        comment.ID,
				"username":   comment.UserName,
				"user_id":    comment.UserID,
				"comment":    comment.Comment,
				"likes":      0,
				"created_at": comment.CreatedAt,
				"updated_at": comment.UpdatedAt,
			},
		},
	}
	_, err := r.db.Collection("news").UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil

}

func (r *repoComment) DeleteComment(newsID *string, commentID *string, ctx context.Context) error {
	filter := bson.M{"_id": *newsID}
	update := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"_id": *commentID,
			},
		},
	}

	result, err := r.db.Collection("news").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("eror ga ada yg kedelete")
	}

	return nil
}

func (r *repoComment) AddLikeComment(newsID, commentID *string, ctx context.Context) error {
	filter := bson.M{"_id": *newsID, "comments._id": *commentID}
	update := bson.M{
		"$inc": bson.M{
			"comments.$.likes": +1,
		},
	}

	rslt, err := r.db.Collection("news").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if rslt.ModifiedCount == 0 {
		return fmt.Errorf("no update")
	}
	return nil
}

func (r *repoComment) GetOneComment(newsID *string, commentID *string, ctx context.Context) (*models.CommentNews, error) {
	filter := bson.M{"_id": *newsID,
		"comments": bson.M{
			"$elemMatch": bson.M{
				"_id": "1",
			},
		},
	}

	result := r.db.Collection("news").FindOne(ctx, filter)

	var comments *models.CommentNews
	err := result.Decode(&comments)
	if err != nil {
		return comments, err
	}

	fmt.Println(comments.Comment)

	return comments, nil
}
