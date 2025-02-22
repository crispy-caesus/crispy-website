package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):] // Extract title from URL path
	if title == "" {
		title = "Home" // Default title if none provided
	}
	p := &Page{Title: title}
	renderTemplate(w, "index", p)
}

// New handler for plain text response
func plainTextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")                   // Set content type to plain text
	w.Write([]byte("dh=b81129b4e9cc388c5ab63919550316fc3ca5ebe4")) // Write the plain text response
}

func main() {
	// Serve static files from the "static" directory at the "/static/" path
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up the main route
	http.HandleFunc("/", viewHandler)

	// Set up a new route for plain text
	http.HandleFunc("/.well-known/discord", plainTextHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

