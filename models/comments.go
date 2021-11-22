package models

import (
	"time"
)

type CommentNews struct {
	ID        string    `json:"comment_id" bson:"_id"`
	UserName  string    `json:"username" bson:"username"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Comment   string    `bson:"comment" json:"comment"`
	Likes     uint      `json:"likes" bson:"likes"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
