package mapper

import (
	"encoding/json"
	"net/http"

	. "../interfaces"
	. "../models"
)

func (m PostMapper) DataMapper(post Post, r *http.Request) Post {
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&post)
	return post
}

type PostMapper struct{}

func NewPostMapper() IPostMapper {
	return PostMapper{}
}
