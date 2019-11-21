package main

import (
	"html/template"
	"net/http"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", home)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}