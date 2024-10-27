package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {}

func NewHandler() *Handler {
  return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
  router.HandleFunc("/login", h.handleLogin).Methods("POST")
  router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello from login route")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello from login route")
}
