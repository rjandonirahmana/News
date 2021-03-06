package models

import (
	"time"
)

type News struct {
	ID        string        `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"title" json:"title"`
	Author    string        `bson:"author" json:"author"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
	Categroy  Category      `bson:"category" json:"category"`
	Location  string        `bson:"location" json:"location"`
	Images    []Images      `bson:"images" json:"images"`
	Comment   []CommentNews `bson:"comments" json:"comments"`
}

type Images struct {
	ImageName string `bson:"image_name" json:"image_name"`
	IsPrimary bool   `bson:"is_primary" json:"is_primary"`
}

type Category struct {
	ID   int
	Name string
}

type CreateNews struct {
	Title      string `bson:"title" json:"title"`
	Author     string `bson:"author" json:"author"`
	Content    string `bson:"content" json:"content"`
	CategroyID int    `bson:"category" json:"category"`
	Location   string `bson:"location" json:"location"`
}
