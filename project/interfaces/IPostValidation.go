package interfaces

import (
	. "../models"
)

type IPostValidation interface {
	PostValidation(post Post) error
}
