package repository

import (
	"context"

	"github.com/rjandonirahmana/news/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminRepo struct {
	db *mongo.Database
}

type AdminRepository interface {
	CreateAdmin(admin *models.Admin, ctx context.Context) error
	GetAdminByID(id *string, ctx context.Context) (*models.Admin, error)
}

func NewAdminRepo(db *mongo.Database) *adminRepo {
	return &adminRepo{db: db}
}

func (r *adminRepo) CreateAdmin(admin *models.Admin, ctx context.Context) error {
	_, err := r.db.Collection("admin").InsertOne(ctx, admin)
	if err != nil {
		return err
	}

	return nil

}

func (r *adminRepo) GetAdminByID(id *string, ctx context.Context) (*models.Admin, error) {
	rslt := r.db.Collection("admin").FindOne(ctx, bson.M{"id": *id})
	var admin *models.Admin

	err := rslt.Decode(&admin)
	if err != nil {
		return admin, err
	}

	return admin, nil
}
