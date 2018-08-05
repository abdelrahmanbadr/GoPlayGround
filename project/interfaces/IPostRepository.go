package interfaces

import (
	. "../models"
)

type IPostRepository interface {
	GetPostById(Id string) (Post, error)
	GetAllPosts() []Post
	RemovePost(id string) ([]Post, error)
	UpdatePost(id string, p Post) (Post, error)
	InsertPost(p Post)
	/////////////
	AddCommentToPost(postId string, c Comment) error
	// RemoveComment(commentId string)
	// UpdateComment(commentId string, c Comment)
	// GetAllPostComments(postId string)
}
