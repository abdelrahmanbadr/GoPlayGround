package repositories

import (
	"errors"

	. "../models"
)

type PostRepository struct {
	Post
}

var posts []Post

func (post *PostRepository) GetPostById(id string) (Post, error) {
	for _, item := range posts {
		if item.Id == id {
			return item, nil
		}
	}
	return Post{}, errors.New("Id Not Found")
}
func (post *PostRepository) GetAllPosts() []Post {
	return posts
}
func (post *PostRepository) InsertPost(p Post) {
	posts = append(posts, p)
}
func (post *PostRepository) Remove(id string) ([]Post, error) {
	for index, item := range posts {
		if item.Id == id {
			posts = append(posts[:index], posts[index+1:]...)
			// break
			return posts, nil
		}
	}
	return posts, errors.New("Id Not Found")
}

//cannot define new method on non-local type model.post
// func (post Post) GetPostById(id string) (Post, error) {
// 	for _, item := range posts {
// 		if item.Id == id {
// 			return item, nil
// 		}
// 	}
// 	return Post{}, errors.New("Id Not Found")
// }

func (post *PostRepository) Update(id string, p Post) (Post, error) {
	for index, item := range posts {
		if item.Id == id {
			posts = append(posts[:index], posts[index+1:]...)
			p.Id = id
			posts = append(posts, p)
			return p, nil
		}
	}
	return Post{}, errors.New("Id Not Found")
}
