package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Dhar01/lenslocked/controllers"
	"github.com/Dhar01/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	// home page
	homeTpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(homeTpl))

	// contact page
	contactTpl := views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	// FAQ page
	faqTpl := views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
