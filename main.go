package main

import "github.com/fourlinecode/golang/filereader"

func main() {
	page := filereader.Page{Title: "A new world", Body: []byte("This is the body for the book.")}

	page.Save()
}
