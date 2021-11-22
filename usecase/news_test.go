package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/rjandonirahmana/news/database"
	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/repository"
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
			news:     models.News{ID: "10", Name: "ngasal", Author: "saya", Content: "assss", Comment: []models.CommentNews{{ID: "1", Comment: "ngakak"}}},
			isTrue:   true,
		},
	}

	db, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	assert.NoError(t, err)

	repo := repository.NewRepoNews(db)
	service := NewServiceNews(repo)

	for _, test := range tetsCases {
		news, err := service.CreateNews(&test.news, context.Background())
		assert.Nil(t, err)

		fmt.Println(news)
	}
}
