package user

import (
	"net/http"

	"Rest-Api/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{} // tip for searching "all works"

const (
	usersURL = "/users/"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetUsersList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Users List")) // empty block. "WIP"
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Get User By UUID")) // empty block. "WIP"
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create User")) // empty block. "WIP"
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Update User")) // empty block. "WIP"
}

func (h *handler) PartUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Partially Update User")) // empty block. "WIP"
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Delete User")) // empty block. "WIP"
}
