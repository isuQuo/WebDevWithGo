package main

import (
	"WebDevWithGo/controllers"
	"WebDevWithGo/views"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(filepath string) views.Template {
	tpl, err := views.Parse(filepath)
	if err != nil {
		panic(err)
	}

	return tpl
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(executeTemplate("templates/home.gohtml")))
	r.Get("/contact", controllers.StaticHandler(executeTemplate("templates/contact.gohtml")))
	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}
