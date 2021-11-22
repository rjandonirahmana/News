package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/rjandonirahmana/news/database"
	"github.com/rjandonirahmana/news/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateNews(t *testing.T) {
	tetsCases := []struct {
		testName string
		news     models.News
		isTrue   bool
	}{
		{
			testName: "1",
			news:     models.News{ID: "5", Name: "testing unit test berhasil", Author: "saya", Content: "assssalamualaikumm", Comment: []models.CommentNews{}, Images: []models.Images{}},
			isTrue:   true,
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoNews(db)

	for _, test := range tetsCases {
		news, err := repo.CreateNews(&test.news, context.Background())
		assert.Nil(t, err)

		fmt.Println(news)
	}
}

func TestUpdateNews(t *testing.T) {
	testCases := []struct {
		testname string
		news     models.News
	}{
		{
			testname: "2",
			news:     models.News{ID: "1", Name: "ganti ini", Author: "ganti orang lain", Content: "berita satu", Comment: []models.CommentNews{{ID: "1", Comment: "ngasal"}}},
		},
	}

	for _, test := range testCases {
		db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
		assert.NoError(t, err)

		repo := NewRepoNews(db)
		newsupdate, err := repo.UpdateNews(&test.news, context.Background())
		assert.Nil(t, err)

		fmt.Println(newsupdate)

	}
}

func TestGetNews(t *testing.T) {
	tetsCases := []struct {
		testName string
		newsID   string
		isTrue   bool
	}{
		{
			testName: "1",
			newsID:   "4",
			isTrue:   true,
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoNews(db)

	for _, test := range tetsCases {
		news, err := repo.GetNewsByID(&test.newsID, context.Background())
		assert.Nil(t, err)

		fmt.Println(news)
	}
}

func TestDeleteNews(t *testing.T) {

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := NewRepoNews(db)

	err = repo.DeleteALLNews(context.Background())
	assert.Nil(t, err)

}
