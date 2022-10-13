package user

import (
	"net/http"

	"Rest-Api/internal/handlers"
	"Rest-Api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{} // tip for searching "all works"

const (
	usersURL = "/users/"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, h.GetUsersList)
	router.HandlerFunc(http.MethodGet, userURL, h.GetUserByUUID)
	router.HandlerFunc(http.MethodPost, userURL, h.CreateUser)
	router.HandlerFunc(http.MethodPut, userURL, h.UpdateUser)
	router.HandlerFunc(http.MethodPatch, userURL, h.PartUpdateUser)
	router.HandlerFunc(http.MethodDelete, userURL, h.DeleteUser)
}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Users List")) // empty block. "WIP"
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Get User By UUID")) // empty block. "WIP"
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("Create User")) // empty block. "WIP"
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Update User")) // empty block. "WIP"
}

func (h *handler) PartUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Partially Update User")) // empty block. "WIP"
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Delete User")) // empty block. "WIP"
}
