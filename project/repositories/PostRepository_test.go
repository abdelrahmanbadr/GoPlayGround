package repositories

import (
	"testing"

	. "../models"
	"github.com/magiconair/properties/assert"
)

func TestPost_Update_Fails(t *testing.T) {
	post := Post{}
	repo := PostRepository{&post}
	post, err := repo.UpdatePost("any_Id", post)
	assert.Equal(t, err.Error(), "Id Not Found")
}

func Test_GetPostById(t *testing.T) {
	p := Post{}
	post := Post{Title: "some title", Content: "Content", Id: "123"}
	repo := PostRepository{&p}
	repo.InsertPost(post)
	var posts []Post

	_ = append(posts, post)
	result, err := repo.GetPostById("123")
	assert.Equal(t, result, post)
	assert.Equal(t, err, nil)
}
