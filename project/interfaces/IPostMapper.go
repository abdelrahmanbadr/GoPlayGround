package interfaces

import (
	"net/http"

	. "../models"
)

type IPostMapper interface {
	DataMapper(post Post, r *http.Request) Post
}
