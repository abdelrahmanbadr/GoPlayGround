package interfaces

import (
	"../models"
)

type IPostRepository interface {
	GetPostById(Id string) (models.Post, error)
}
