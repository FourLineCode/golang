package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fourlinecode/golang/filereader"
)

type httpResponse struct {
	title string
	body  string
}

// RootHandler for /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Wiki Api")
}

// PageHandler for /page
func PageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/page/"):]
	if title == "" {
		http.Error(w, "Title not provided", 400)
		return
	}

	page, err := filereader.Load(title)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	res := page.Body

	json.NewEncoder(w).Encode(res)
}
