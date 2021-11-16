package models

import (
	"time"
)

type News struct {
	ID         string        `bson:"id" json:"id"`
	Name       string        `bson:"title" json:"title"`
	Author     string        `bson:"author" json:"author"`
	Content    string        `bson:"content" json:"content"`
	Created_at time.Time     `bson:"created_at" json:"created_at"`
	Updated_at time.Time     `bson:"updated_at" json:"updated_at"`
	Images     []Images      `bson:"images" json:"images"`
	Comment    []CommentNews `bson:"comments" json:"comments"`
}

type Images struct {
	ImageName string `bson:"image_name" json:"image_name"`
	IsPrimary bool   `bson:"is_primary" json:"is_primary"`
}
