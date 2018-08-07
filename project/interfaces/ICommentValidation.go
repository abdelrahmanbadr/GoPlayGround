package interfaces

import (
	. "../models"
)

type ICommentValidation interface {
	CommentValidation(comment Comment) error
}
