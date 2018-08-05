package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListPosts(t *testing.T) {

	req, _ := http.NewRequest("GET", "", nil)

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(ListPosts)

	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
