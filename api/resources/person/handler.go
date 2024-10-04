package person

import "net/http"

type Handler interface {
	ListPeople(w http.ResponseWriter, r *http.Request)
	GetPerson(w http.ResponseWriter, r *http.Request)
	CreatePerson(w http.ResponseWriter, r *http.Request)
	UpdatePerson(w http.ResponseWriter, r *http.Request)
	DeletePerson(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) ListPeople(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
