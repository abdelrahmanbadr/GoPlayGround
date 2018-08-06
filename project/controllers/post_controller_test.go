package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListPosts(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/api/posts", nil)
	if err != nil {
		t.Fatalf("couldn't create request")
	}

	resp := httptest.NewRecorder()

	hf := http.HandlerFunc(ListPosts)

	hf.ServeHTTP(resp, req)

	// // Check the status code is what we expect.
	checkResponseCode(t, http.StatusOK, resp.Code)

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentPost(t *testing.T) {

	req, _ := http.NewRequest("GET", "localhost:8080/api/posts/11", nil)
	// response := executeRequest(req)
	resp := httptest.NewRecorder()

	hf := http.HandlerFunc(GetPost)

	hf.ServeHTTP(resp, req)

	checkResponseCode(t, http.StatusNotFound, resp.Code)

	var m map[string]string
	json.Unmarshal(resp.Body.Bytes(), &m)
	// fmt.Println(m)
	if m["error"] != "Id Not Found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Id Not Found'. Got '%s'", m["error"])
	}
}
