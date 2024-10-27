package main

import (
	"log"

	"github.com/dtg-lucifer/go-backend/cmd/api"
)

func main() {
  server := api.NewApiServer(":8080", nil)
  if err := server.Run(); err != nil {
    log.Fatal("Error: Running the server", err)
  }
}
