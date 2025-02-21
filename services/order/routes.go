package order

import (
	"net/http"

	"github.com/dtg-lucifer/go-backend/typedef"
	"github.com/gorilla/mux"
)

type Handler struct {
  store typedef.OrderStore
}

func NewHandler(store typedef.OrderStore) *Handler {
  return &Handler{ store: store }
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
  router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}
