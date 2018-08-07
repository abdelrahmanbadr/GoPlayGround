package validation

import (
	. "../interfaces"
	. "../models"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type PostValidation struct{}

func NewPostValidation() IPostValidation {
	return PostValidation{}
}

func PostStructLevelValidation(sl validator.StructLevel) {

	post := sl.Current().Interface().(Post)

	if len(post.Title) == 0 {
		sl.ReportError(post.Title, "Title", "title", "not valid, should be string and it's required", "")
	}
	if len(post.Content) == 0 {
		sl.ReportError(post.Content, "Content", "content", "not valid, should be string and it's required", "")
	}
}

func (v PostValidation) PostValidation(post Post) error {
	validate = validator.New()
	validate.RegisterStructValidation(PostStructLevelValidation, Post{})
	err := validate.Struct(post)
	return err
}
