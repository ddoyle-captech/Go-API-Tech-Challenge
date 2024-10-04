package router

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"Go-API-Tech-Challenge/api/resources/person"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Creates a concrete implementation of chi.Router. Registers all API
// endpoints and adds any middleware.
func New() chi.Router {
	r := chi.NewRouter()

	ph := person.NewHandler()
	ch := course.NewHandler()

	r.Route("/api/person", func(r chi.Router) {
		r.Get("/", http.HandlerFunc(ph.ListPeople))
		r.Get("/{name}", http.HandlerFunc(ph.GetPerson))
		r.Post("/", http.HandlerFunc(ph.CreatePerson))
		r.Put("/{name}", http.HandlerFunc(ph.UpdatePerson))
		r.Delete("/{name}", http.HandlerFunc(ph.DeletePerson))
	})

	r.Route("/api/course", func(r chi.Router) {
		r.Get("/", http.HandlerFunc(ch.ListCourses))
		r.Get("/{id}", http.HandlerFunc(ch.GetCourse))
		r.Post("/", http.HandlerFunc(ch.CreateCourse))
		r.Put("/{id}", http.HandlerFunc(ch.UpdateCourse))
		r.Delete("/{id}", http.HandlerFunc(ch.DeleteCourse))
	})

	return r
}
