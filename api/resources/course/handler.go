package course

import (
	"Go-API-Tech-Challenge/api"
	"encoding/json"
	"log"
	"net/http"
)

// Handler is responsible for all HTTP communication for the /api/course endpoints. It
// manages parsing params and serializing responses to return to callers.
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
		log.Printf("unable to fetch courses from database, error: %s\n", err.Error())
		resp := api.ErrorResponse{
			Message: "unable to fetch courses",
		}
		resp.Send(w, http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(courses)
	if err != nil {
		log.Printf("unable to serialize courses to JSON, error: %s\n", err.Error())
		resp := api.ErrorResponse{
			Message: "unable to fetch courses",
		}
		resp.Send(w, http.StatusInternalServerError)
		return
	}

	w.Write(body)
	w.Header().Set("Content-Type", "application/json")
}

func (h *handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")

	// w.Write(body)
	w.Header().Set("Content-Type", "application/json")
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
