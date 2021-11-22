package repository

import (
	"context"
	"testing"

	"github.com/rjandonirahmana/news/database"
	"github.com/rjandonirahmana/news/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	tetsCases := []struct {
		newsID  string
		comment models.CommentNews
		isTrue  bool
	}{
		{
			newsID:  "5",
			comment: models.CommentNews{ID: "1", UserName: "atengbaruangett", Comment: "apalah kauu kereen kali wak kau"},
			isTrue:  true,
		}, {
			newsID:  "4",
			comment: models.CommentNews{ID: "2", UserName: "atengbaru", Comment: "apalah apalah25", UserID: "1"},
			isTrue:  true,
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoComments(db)

	for _, test := range tetsCases {
		err := repo.CreateComment(&test.newsID, &test.comment, context.Background())
		assert.Nil(t, err)
	}
}

func TestDeleteComment(t *testing.T) {
	tetsCases := []struct {
		testName  string
		newsID    string
		commentID string
		isTrue    bool
	}{
		{
			testName:  "3",
			newsID:    "4",
			commentID: "1",
			isTrue:    true,
		}, {
			testName:  "3",
			newsID:    "4",
			commentID: "2",
			isTrue:    true,
		}, {
			testName:  "3",
			newsID:    "4",
			commentID: "3",
			isTrue:    true,
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoComments(db)

	for _, test := range tetsCases {
		err := repo.DeleteComment(&test.newsID, &test.commentID, context.TODO())
		assert.Nil(t, err)
	}
}

func TestLikeComment(t *testing.T) {
	tetsCases := []struct {
		newsID    string
		commentID string
	}{
		{
			newsID:    "4",
			commentID: "1",
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoComments(db)

	for _, test := range tetsCases {
		err := repo.AddLikeComment(&test.newsID, &test.commentID, context.Background())
		assert.Nil(t, err)
	}
}

func TestUpdateComment(t *testing.T) {
	tetsCases := []struct {
		newsID  string
		comment models.CommentNews
	}{
		{
			newsID:  "4",
			comment: models.CommentNews{ID: "1", Comment: "ini update comment di unit test kali kedua nyaaa bisaaa apalah"},
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoComments(db)

	for _, test := range tetsCases {
		err := repo.UpdateComment(&test.newsID, &test.comment, context.Background())
		assert.Nil(t, err)
	}
}
