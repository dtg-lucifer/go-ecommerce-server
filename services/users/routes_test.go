package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/dtg-lucifer/go-backend/typedef"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := typedef.RegisterUserPaylod{
			FirstName: "Piush",
			LastName:  "Bose",
			Email:     "invalid",
			Password:  "6789",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should succeed if the user payload is valid", func(t *testing.T) {
		payload := typedef.RegisterUserPaylod{
			FirstName: "Piush",
			LastName:  "Bose",
			Email:     "piush@gmail.com",
			Password:  "123456789",
		}

		marshalled, _ := json.Marshal(payload)
		// log.Print(bytes.NewBuffer(marshalled))
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByID(id int) (*typedef.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*typedef.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user typedef.User) error {
	return nil
}
