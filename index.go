package main

import (
        "log"
        "net/http"
        "html/template"
)

type Page struct {
        Title string
        Body []byte
}

var templates = template.Must(template.ParseFiles("index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
        err := templates.ExecuteTemplate(w, tmpl+".html", p)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
        p := &Page{Title: "index"}
        renderTemplate(w, "index", p)
}


func main() {
        http.HandleFunc("/", viewHandler)
        log.Fatal(http.ListenAndServe(":8080", nil))
}

