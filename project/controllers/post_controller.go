package controllers

import (
	"encoding/json"
	"net/http"

	. "../common"
	. "../mapper"
	. "../models"
	. "../repositories"
	. "../validation"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
var RepositoryInstance = PostRepository{}
var MapperInstance = PostMapper{}
var ValidationInstance = PostValidation{}

func AddPost(w http.ResponseWriter, r *http.Request) {
	// it should be separated in other file or it's fine ??
	validate = validator.New()
	validate.RegisterStructValidation(ValidationInstance.PostStructLevelValidation, Post{})
	post := NewPost()
	post = MapperInstance.DataMapper(post, r)
	err := validate.Struct(post)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RepositoryInstance.InsertPost(post)
	json.NewEncoder(w).Encode(post)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	posts := RepositoryInstance.GetAllPosts()
	json.NewEncoder(w).Encode(posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	posts, err := RepositoryInstance.RemovePost(postId)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	post, err := RepositoryInstance.GetPostById(postId)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	var post = Post{}
	post = MapperInstance.DataMapper(post, r)
	post, err := RepositoryInstance.UpdatePost(postId, post)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, post)
}

func AddComment(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	params := mux.Vars(r)
	postId := params["id"]
	comment := NewComment()
	json.NewDecoder(r.Body).Decode(&comment)
	err := RepositoryInstance.AddCommentToPost(postId, comment)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	json.NewEncoder(w).Encode(comment)

}
