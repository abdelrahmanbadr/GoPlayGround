package interfaces

import (
	"net/http"

	. "../models"
)

type ICommentMapper interface {
	DataMapper(comment Comment, r *http.Request) Comment
}
