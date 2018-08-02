package controllers

import (
	"encoding/json"
	"net/http"

	. "../models"
	"github.com/gorilla/mux"
)

var PostInstance = Post{}

func AddPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	post := NewPost()
	json.NewDecoder(r.Body).Decode(&post)
	post.InsertPost(post)
	json.NewEncoder(w).Encode(post)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	posts := PostInstance.GetAllPosts()
	json.NewEncoder(w).Encode(posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	posts := PostInstance.Remove(postId)
	json.NewEncoder(w).Encode(posts)

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	post := PostInstance.GetPostById(postId)
	json.NewEncoder(w).Encode(post)
	// if no id exists will return empty object
	// json.NewEncoder(w).Encode(&models.Book{})

}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	postId := params["id"]
	var post = PostInstance
	json.NewDecoder(r.Body).Decode(&post)
	post = PostInstance.Update(postId, post)
	json.NewEncoder(w).Encode(post)

}
