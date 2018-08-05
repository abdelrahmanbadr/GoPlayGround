package models

import (
	"github.com/twinj/uuid"
)

type Post struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"auther"`
	Comments []Comment `json:"comments"`
	// UpdatedAt time.Time `json:"updated_at"`
	// CreatedAt time.Time `json:"created_at"`
}

func NewPost() Post {
	var row Post
	row.Id = uuid.NewV4().String()
	return row
}
