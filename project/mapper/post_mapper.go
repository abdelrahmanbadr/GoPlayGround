package mapper

import (
	"encoding/json"
	"net/http"

	. "../models"
)

func DataMapper(r *http.Request) Post {
	defer r.Body.Close()
	post := NewPost()
	json.NewDecoder(r.Body).Decode(&post)
	return post
}
