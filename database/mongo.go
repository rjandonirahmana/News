package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionMongo(host string, dbname string, ctx context.Context) (*mongo.Database, error) {

	clientoption := options.Client()
	clientoption.ApplyURI(host)

	client, err := mongo.NewClient(clientoption.SetMaxConnIdleTime(time.Duration(50) * time.Second))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbname)
	return db, nil
}
