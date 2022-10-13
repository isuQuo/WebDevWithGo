package main

import (
	"WebDevWithGo/views"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := views.Parse(filepath)
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
	//w.WriteHeader(http.StatusNotFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}
