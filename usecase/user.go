package usecase

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type usecaseUser struct {
	repo      repository.RepoUser
	repoadmin repository.AdminRepository
	secret    string
}

type UsecaseUser interface {
	CreateUser(user *models.User, ctx context.Context) (*models.User, error)
	LoginUser(email, password *string, ctx context.Context) (*models.User, string, error)
	GetUserByID(id *string, ctx context.Context) (*models.User, error)
}

func NewUsecCaseUser(repo repository.RepoUser, secret string, repoadmin repository.AdminRepository) *usecaseUser {
	return &usecaseUser{repo: repo, secret: secret, repoadmin: repoadmin}
}

func RandRuneSalt(n uint) string {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
}

func (u *usecaseUser) CreateUser(user *models.User, ctx context.Context) (*models.User, error) {
	err := u.repo.IsEmailAvailable(&user.Email, ctx)
	if err != nil {
		return user, err
	}

	user.ID = uuid.New().String()
	user.Salt = RandRuneSalt(10)
	h := sha256.New()
	h.Write([]byte(user.Password + user.Salt))
	user.Password = fmt.Sprintf("%x", h.Sum([]byte(u.secret)))
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	err = u.repo.CreateUser(user, ctx)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (u *usecaseUser) LoginUser(email, password *string, ctx context.Context) (*models.User, string, error) {
	user, err := u.repo.GetUserByEmail(email, ctx)
	if err != nil {
		return user, "", err
	}

	h := sha256.New()
	h.Write([]byte(*password + user.Salt))
	*password = fmt.Sprintf("%x", h.Sum([]byte(u.secret)))
	if *password != user.Password {
		return nil, "", fmt.Errorf("please input your password correctly")
	}

	var roles string
	_, err = u.repoadmin.GetAdminByID(&user.ID, ctx)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, "", err
	}

	if err == mongo.ErrNoDocuments {
		roles = "user"
	} else {
		roles = "admin"
	}

	return user, roles, nil
}

func (u *usecaseUser) GetUserByID(id *string, ctx context.Context) (*models.User, error) {
	return u.repo.GetUserByID(id, ctx)
}
