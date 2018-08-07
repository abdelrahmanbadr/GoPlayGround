package mapper

import (
	"encoding/json"
	"net/http"

	. "../interfaces"
	. "../models"
)

func (m CommentMapper) DataMapper(comment Comment, r *http.Request) Comment {
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&comment)
	return comment
}

type CommentMapper struct{}

func NewCommentMapper() ICommentMapper {
	return CommentMapper{}
}
