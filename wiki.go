package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// This method's signature reads: "This is a method named save that takes as its receiver p,
// a pointer to Page . It takes no parameters, and returns a value of type error."
//
// The octal integer literal 0600, passed as the third parameter to WriteFile,
// indicates that the file should be created with read-write permissions for the current user only.
// (See the Unix man page open(2) for details.)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[(len("/view/")):]
	p, _ := loadPage(title)
	fmt.Fprint(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

const Port = "8080"

func main() {
	http.HandleFunc("/view/", viewHandler)
	addr := fmt.Sprintf(":%s", Port)

	println("Listening on port: " + Port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	} else {
	}
}
