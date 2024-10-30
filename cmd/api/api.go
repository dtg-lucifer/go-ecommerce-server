package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dtg-lucifer/go-backend/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
  addr string
  db *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
  return &APIServer{
    addr: addr,
    db: db,
  }
}

func (s *APIServer) Run() error {
  router := mux.NewRouter()
  subrouter := router.PathPrefix("/api/v1").Subrouter()

  store := users.NewStore(s.db)
  userHandler := users.NewHandler(store)
  userHandler.RegisterRoutes(subrouter)

  log.Println("Listening on", s.addr) 

  return http.ListenAndServe(s.addr, router)  
}
