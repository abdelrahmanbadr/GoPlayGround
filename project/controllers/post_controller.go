package controllers

import (
	"encoding/json"
	"net/http"

	. "../common"
	. "../models"
	. "../repositories"
	"github.com/gorilla/mux"
)

var RepositoryInstance = PostRepository{}

func AddPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	post := NewPost()
	json.NewDecoder(r.Body).Decode(&post)
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
	defer r.Body.Close()
	params := mux.Vars(r)
	postId := params["id"]
	var post = Post{}
	json.NewDecoder(r.Body).Decode(&post)
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
