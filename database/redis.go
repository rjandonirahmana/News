package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisConnection(username, password string) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: username,
		Password: password,
		DB:       5,
		PoolSize: 20,
	})

	_, err := db.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Println("Redis is connected...")

	return db, nil
}
