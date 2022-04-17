package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/go-redis/redis/v8"
)

func randStringBytes(s int) (string, error) {
	b := make([]byte, s)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

type Authredis struct {
	rdb *redis.Client
}

type AuthToken struct {
	UserID string `redis:"user_id"`
	Roles  string `redis:"roles"`
	Token  string `redis:"token"`
}

func NewAuthRedis(redis *redis.Client) *Authredis {
	return &Authredis{rdb: redis}
}

type Authentication interface {
	GenerateToken(ctx context.Context, email, userID, roles string) (string, error)
	GetTokenRedis(ctx context.Context, token string) (AuthToken, error)
}

func (r *Authredis) GenerateToken(ctx context.Context, email, userID, roles string) (string, error) {

	token, err := randStringBytes(25)
	if err != nil {
		return token, err
	}

	exist, err := r.rdb.Exists(ctx, email).Result()
	if err != nil {
		return "", err
	}
	if exist == 1 {
		oldtoken, err := r.rdb.HGet(ctx, email, "token").Result()
		if err != nil {
			return "", err
		}
		_, err = r.rdb.Del(ctx, oldtoken).Result()
		if err != nil {
			return "", err
		}
		pipe := r.rdb.TxPipeline()
		pipe.HSet(ctx, email, "token", token)
		pipe.HSet(ctx, email, "user_id", userID)
		pipe.HSet(ctx, email, "roles", roles)
		pipe.Set(ctx, token, email, time.Hour*24)

		_, err = pipe.Exec(ctx)
		if err != nil {
			return "", err
		}
	} else {
		pipe := r.rdb.TxPipeline()
		pipe.HSet(ctx, email, "user_id", userID, "roles", roles, "token", token)
		pipe.Set(ctx, token, email, time.Hour*24)
		_, err := pipe.Exec(ctx)

		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func (r *Authredis) GetTokenRedis(ctx context.Context, token string) (AuthToken, error) {
	var auth AuthToken
	email, err := r.rdb.Get(ctx, token).Result()
	if err != nil {
		return auth, err
	}

	err = r.rdb.HGetAll(ctx, email).Scan(&auth)

	if err != nil {
		return auth, err
	}

	return auth, nil
}
