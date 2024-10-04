package course

import "net/http"

type Handler interface {
	ListCourses(w http.ResponseWriter, r *http.Request)
	GetCourse(w http.ResponseWriter, r *http.Request)
	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	DeleteCourse(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) ListCourses(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
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
