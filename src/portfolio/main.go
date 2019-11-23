package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	images := http.FileServer(http.Dir("images"))
	mux.Handle("/images/", http.StripPrefix("/images/", images))
	mux.HandleFunc("/", home)
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}
