package validation

import (
	. "../interfaces"
	. "../models"
	"gopkg.in/go-playground/validator.v9"
)

type PostValidation struct{}

func NewPostValidation() IPostValidation {
	return PostValidation{}
}

func (v PostValidation) PostStructLevelValidation(sl validator.StructLevel) {

	post := sl.Current().Interface().(Post)

	if len(post.Title) == 0 {
		sl.ReportError(post.Title, "Title", "title", "not valid, should be string and it's required", "")
	}
	if len(post.Content) == 0 {
		sl.ReportError(post.Content, "Content", "content", "not valid, should be string and it's required", "")
	}
}
