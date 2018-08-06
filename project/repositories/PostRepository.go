package repositories

import (
	"errors"

	. "../interfaces"
	. "../models"
)

type PostRepository struct {
	post *Post
}

func NewPostRepository(post *Post) IPostRepository {
	return &PostRepository{post}
}

var posts []Post

func (p *PostRepository) GetPostById(id string) (Post, error) {
	for _, item := range posts {
		if item.Id == id {
			return item, nil
		}
	}
	return Post{}, errors.New("Id Not Found")
}
func (p *PostRepository) GetAllPosts() []Post {
	return posts
}
func (p *PostRepository) InsertPost(post Post) {
	posts = append(posts, post)
}
func (p *PostRepository) RemovePost(id string) ([]Post, error) {
	for index, item := range posts {
		if item.Id == id {
			posts = append(posts[:index], posts[index+1:]...)
			// break
			return posts, nil
		}
	}
	return posts, errors.New("Id Not Found")
}

func (p *PostRepository) UpdatePost(id string, post Post) (Post, error) {
	for index, item := range posts {
		if item.Id == id {
			post = posts[index]
			post.Id = id
			posts = append(posts, post)
			return post, nil
		}
	}
	return Post{}, errors.New("Id Not Found")
}

func (p *PostRepository) AddCommentToPost(postId string, c Comment) error {
	post, err := p.GetPostById(postId)
	if err != nil {
		return err
	}
	post.Comments = append(post.Comments, c)
	_, err = p.UpdatePost(postId, post)
	if err != nil {
		return err
	}
	return nil
}
