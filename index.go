package main

import (
	"log"
	"net/http"
)


func viewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func plainTextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("dh=b81129b4e9cc388c5ab63919550316fc3ca5ebe4"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/music", musicHandler)

	http.HandleFunc("/.well-known/discord", plainTextHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

