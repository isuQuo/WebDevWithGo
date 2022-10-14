package main

import (
	"WebDevWithGo/controllers"
	"WebDevWithGo/templates"
	"WebDevWithGo/views"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(filepath ...string) views.Template {
	tpl, err := views.ParseFS(templates.FS, filepath...)
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
	r.Get("/", controllers.StaticHandler(executeTemplate("home.gohtml", "tailwind.gohtml")))
	r.Get("/contact", controllers.StaticHandler(executeTemplate("contact.gohtml", "tailwind.gohtml")))
	r.Get("/faq", controllers.FAQ(executeTemplate("faq.gohtml", "tailwind.gohtml")))
	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}
