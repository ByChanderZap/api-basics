package user

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ByChanderZap/api-basics/types"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test",
			LastName:  "last",
			Email:     "invalid",
			Password:  "asd",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := chi.NewMux()
		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should register a new user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test",
			LastName:  "last",
			Email:     "mail@test.com",
			Password:  "asd",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := chi.NewMux()
		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var createdUser User
		if err := json.NewDecoder(rr.Body).Decode(&createdUser); err != nil {
			t.Fatal(err)
		}

		if createdUser.Email != payload.Email {
			t.Errorf("expected email %s, got %s", payload.Email, createdUser.Email)
		}
		if createdUser.FirstName != payload.FirstName {
			t.Errorf("expected first name %s, got %s", payload.FirstName, createdUser.FirstName)
		}
		if createdUser.LastName != payload.LastName {
			t.Errorf("expected last name %s, got %s", payload.LastName, createdUser.LastName)
		}
	})
}

type mockUserStore struct {
}

func (um mockUserStore) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	return User{
		ID:        uuid.New(),
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Email:     arg.Email,
		Password:  arg.Password,
		CreatedAt: arg.CreatedAt,
		UpdatedAt: arg.UpdatedAt,
	}, nil
}

func (um mockUserStore) GetUserByEmail(ctx context.Context, email string) (User, error) {
	if email == "mail@test.com" {
		return User{}, sql.ErrNoRows
	}
	return User{}, fmt.Errorf("user not found")
}

func (um mockUserStore) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	return User{}, nil
}
