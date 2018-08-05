package models

import (
	"github.com/twinj/uuid"
)

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content" `
	// UpdatedAt time.Time `json:"updated_at"`
	// CreatedAt time.Time `json:"created_at"`
}

func NewComment() Comment {
	var row Comment
	row.Id = uuid.NewV4().String()
	return row
}
