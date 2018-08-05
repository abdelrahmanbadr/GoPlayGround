package main

import (
	"log"
	"net/http"

	. "./controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/posts", ListPosts).Methods("GET")
	router.HandleFunc("/api/posts", AddPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", DeletePost).Methods("DELETE")
	router.HandleFunc("/api/posts/{id}", GetPost).Methods("GET")
	router.HandleFunc("/api/posts/{id}", UpdatePost).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
