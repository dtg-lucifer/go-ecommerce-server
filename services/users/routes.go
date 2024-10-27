package users

import (
	"net/http"

	"github.com/dtg-lucifer/go-backend/typedef"
	"github.com/dtg-lucifer/go-backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload typedef.RegisterUserPaylod

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
	}

  // ? Check if the user exists
  
}
