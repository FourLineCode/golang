package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fourlinecode/golang/handlers"
)

func main() {
	// http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/page/", handlers.PageHandler)

	err := http.ListenAndServe(":2333", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on port :2333")
}
