package services

import (
	. "../models"
)

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
			p.Id = id
			posts = append(posts, p)
			return p
		}
	}
	return Post{}
}
