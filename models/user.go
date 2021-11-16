package models

import (
	"time"
)

type User struct {
	UserName   string    `bson:"username" json:"username"`
	ID         string    `bson:"id" json:"user_id"`
	Email      string    `bson:"email" json:"email"`
	Salt       string    `bson:"salt" json:"-"`
	Password   string    `bson:"password" json:"password"`
	Created_at time.Time `bson:"created_at" json:"created_at"`
	Updated_at time.Time `bson:"updated_at" json:"updated_at"`
}
