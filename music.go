package main

import (
	"net/http"
	"text/template"
)

func musicHandler(w http.ResponseWriter, r *http.Request) {
	results := loadMusic()
	tmpl := template.Must(template.ParseFiles("templates/music.html"))

	err := tmpl.Execute(w, results)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
