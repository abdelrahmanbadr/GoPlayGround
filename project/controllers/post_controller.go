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
	posts, err := PostInstance.Remove(postId)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	post, err := PostInstance.GetPostById(postId)
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
	var post = PostInstance
	json.NewDecoder(r.Body).Decode(&post)
	post, err := PostInstance.Update(postId, post)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, post)

}

// separate them
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
