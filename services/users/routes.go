package users

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/dtg-lucifer/go-backend/services/auth"
	"github.com/dtg-lucifer/go-backend/typedef"
	"github.com/dtg-lucifer/go-backend/utils"
)

type Handler struct {
	store typedef.UserStore
}

func NewHandler(store typedef.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	var payload typedef.LoginUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
	}

	if err := utils.BodyValidator.Struct(&payload); err != nil { 							//? Validate request body fields
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, error)
		return
	}

	
	_, err := h.store.GetUserByEmail(payload.Email) 													//? Check if the user exists
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email {%s} already extists", payload.Email))
	}

	hashedPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(typedef.User{
		Email:    payload.Email,
		Password: hashedPass,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, typedef.RegisterResponse{
		Message: "Successfully created the user!",
		Data: typedef.User{
			Email:    payload.Email,
			Password: payload.Password,
		},
	})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload typedef.RegisterUserPaylod

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
	}

	if err := utils.BodyValidator.Struct(&payload); err != nil {								//? Validate request body fields
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, error)
		return
	}
	
	_, err := h.store.GetUserByEmail(payload.Email)															//? Check if the user exists
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email {%s} already extists", payload.Email))
	}

	hashedPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(typedef.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPass,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, typedef.RegisterResponse{
		Message: "Successfully created the user!",
		Data: typedef.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  payload.Password,
		},
	})
}
