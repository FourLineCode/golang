package main

import (
	"fmt"
	"log"

	"github.com/fourlinecode/golang/filereader"
)

func main() {
	page := filereader.Page{Title: "A new world", Body: []byte("This is the body for the book.")}

	page.Save()

	newPage, err := filereader.Load("Hello world")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title: %s\nBody: %s", newPage.Title, newPage.Body)
}
