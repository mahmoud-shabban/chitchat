package main

import (
	"html/template"
	"net/http"

	"github.com/mahmoud-shabban/chitchat/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}

	templates := template.Must(template.ParseFiles(files...))

	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
func main() {

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// server static files
	files := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	server.ListenAndServe()
}
