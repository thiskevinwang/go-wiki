package main

import (
	"fmt"
	"html/template"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[(len("/view/")):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "templates/view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "templates/edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

const Port = "8080"

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	addr := fmt.Sprintf(":%s", Port)

	println(fmt.Sprintf("🚀 Ready on http://localhost:%s", Port))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	} else {

	}
}
