package models

import (
	"errors"

	"github.com/twinj/uuid"
)

var posts []Post

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"auther"`
	// UpdatedAt time.Time `json:"updated_at"`
	// CreatedAt time.Time `json:"created_at"`
}

func NewPost() Post {
	var row Post
	row.Id = uuid.NewV4().String()
	return row
}

// create post
func (post Post) InsertPost(p Post) {
	posts = append(posts, p)
}

func (post Post) GetAllPosts() []Post {
	return posts
}
func (post Post) Remove(id string) ([]Post, error) {
	for index, item := range posts {
		if item.Id == id {
			posts = append(posts[:index], posts[index+1:]...)
			// break
			return posts, nil
		}
	}
	return posts, errors.New("Id Not Found")
}
func (post Post) GetPostById(id string) (Post, error) {
	for _, item := range posts {
		if item.Id == id {
			return item, nil
		}
	}
	return Post{}, errors.New("Id Not Found")
}

func (post Post) Update(id string, p Post) (Post, error) {
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
