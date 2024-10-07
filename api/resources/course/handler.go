package course

import (
	"Go-API-Tech-Challenge/api"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	idParam := chi.URLParam(r, "id")

	// Convert id URL parameter to string. If its an invalid integer, return 400 Bad Request
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("received invalid course ID, error: %s\n", err.Error())
		resp := api.ErrorResponse{
			Message: fmt.Sprintf("course ID: %s is invalid", idParam),
		}
		resp.Send(w, http.StatusBadRequest)
		return
	}

	c, err := h.r.FetchCourseByID(id)

	// If no course is found for the given ID, return a 404 Not Found
	if err != nil && errors.Is(err, ErrCourseNotFound) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// If the repo returns an error, return a 500
	if err != nil {
		log.Printf("unexpected error fetching course with ID: %d, error: %s\n", id, err.Error())
		resp := api.ErrorResponse{
			Message: fmt.Sprintf("unable to fetch course with ID: %s", idParam),
		}
		resp.Send(w, http.StatusInternalServerError)
		return
	}

	// Serialize course to JSON
	body, err := json.Marshal(c)
	if err != nil {
		log.Printf("unable to serialize course to JSON, error: %s\n", err.Error())
		resp := api.ErrorResponse{
			Message: fmt.Sprintf("unable to fetch course with ID: %d", id),
		}
		resp.Send(w, http.StatusInternalServerError)
		return
	}

	w.Write(body)
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
