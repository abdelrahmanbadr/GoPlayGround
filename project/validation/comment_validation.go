package validation

import (
	. "../interfaces"
	. "../models"
	"gopkg.in/go-playground/validator.v9"
)

type CommentValidation struct{}

func NewCommentValidation() ICommentValidation {
	return CommentValidation{}
}

func CommentStructLevelValidation(sl validator.StructLevel) {

	comment := sl.Current().Interface().(Comment)

	if len(comment.Content) == 0 {
		sl.ReportError(comment.Content, "Content", "content", "not valid, should be string and it's required", "")
	}
}

func (v CommentValidation) CommentValidation(comment Comment) error {
	validate = validator.New()
	validate.RegisterStructValidation(CommentStructLevelValidation, Comment{})
	err := validate.Struct(comment)
	return err
}
