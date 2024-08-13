package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ebosetalee/password-service.git/types"
	"github.com/gorilla/mux"
)

type mockUserRepo struct {
}

func (m *mockUserRepo) GetUserByEmail(email string) (*types.User, error) {

	return nil, fmt.Errorf("missing request")
}

func (m *mockUserRepo) CreateUser(user types.User) error {

	return nil
}

func TestUser(t *testing.T) {
	userRepo := &mockUserRepo{}
	handler := NewHandler(userRepo)

	payload := types.RegisterPayload{
		FirstName: "John Doe",
		LastName:  "John Doe",
		Email:     "",
		Username:  "JD",
		Password:  "password",
	}

	res, _ := json.Marshal(payload)

	t.Run("Should fail if payload is empty or invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(res))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusPreconditionFailed {
			t.Errorf("expected status code %d, got status code %d", http.StatusPreconditionFailed, rr.Code)
		}
	})
}
