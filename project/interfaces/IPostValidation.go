package interfaces

import (
	"gopkg.in/go-playground/validator.v9"
)

type IPostValidation interface {
	PostStructLevelValidation(sl validator.StructLevel)
}
