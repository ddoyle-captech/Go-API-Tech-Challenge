package course

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler interface {
	ListCourses(w http.ResponseWriter, r *http.Request)
	GetCourse(w http.ResponseWriter, r *http.Request)
	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	DeleteCourse(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	r Repository
}

func NewHandler(r Repository) Handler {
	return &handler{
		r: r,
	}
}

func (h *handler) ListCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.r.FetchCourses()
	if err != nil {
		// TODO: return different responses depending on the error returned
		panic("not implemented")
	}

	body, err := json.Marshal(courses)
	if err != nil {
		log.Printf("unable to serialize courses to JSON, error: %s\n", err.Error())
		http.Error(w, "unable to fetch courses", http.StatusInternalServerError)
	}

	w.Write(body)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
