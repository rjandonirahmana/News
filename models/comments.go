package models

import (
	"time"
)

type CommentNews struct {
	UserName  string    `bson:"username" json:"username"`
	UserID    string    `bson:"user_id" json:"user_id"`
	Comment   string    `bson:"comment" json:"comment"`
	Likes     uint      `bson:"likes" json:"likes" validate:"omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
