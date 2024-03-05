package main

import (
	"fmt"
	"net/http"

	"github.com/Dhar01/lenslocked/controllers"
	"github.com/Dhar01/lenslocked/templates"
	"github.com/Dhar01/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	// home page
	// homeTpl := views.Must(views.ParseFs(templates.FS, "home.gohtml", "layout-parts.gohtml"))
	homeTpl := views.Must(views.ParseFs(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(homeTpl))

	// contact page
	// contactTpl := views.Must(views.ParseFs(templates.FS, "contact.gohtml"))
	contactTpl := views.Must(views.ParseFs(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	// FAQ page

	faqTpl := views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	// r.Get("/faq", controllers.FAQ(views.Must(views.ParseFs(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
