package models

import (
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
func (post Post) Remove(id string) []Post {
	for index, item := range posts {
		if item.Id == id {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	return posts
}
func (post Post) GetPostById(id string) Post {
	for _, item := range posts {
		if item.Id == id {
			return item
		}
	}
	return Post{}
}
func (post Post) Update(id string, p Post) Post {
	for index, item := range posts {
		if item.Id == id {
			posts = append(posts[:index], posts[index+1:]...)
			posts = append(posts, p)
			return p
		}
	}
	return Post{}
}
