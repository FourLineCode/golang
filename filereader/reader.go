package filereader

import "io/ioutil"

// Page of a book
type Page struct {
	Title string `json:"title"`
	Body  []byte `json:"body"`
}

// Save the book
func (p *Page) Save() error {
	filename := p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load a page
func Load(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
