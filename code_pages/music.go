package code_pages

import (
	"net/http"
	"text/template"
)

type musicData struct {
	Releases []musicItem
	Spending float32
}

func MusicHandler(w http.ResponseWriter, r *http.Request) {
	results := loadMusic()
	spending := getSpending()
	tmpl := template.Must(template.ParseFiles("templates/music.html"))

	data := musicData{Releases: results, Spending: spending}

	err := tmpl.Execute(w, data)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
