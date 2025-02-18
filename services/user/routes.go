package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ByChanderZap/api-basics/services/auth"
	"github.com/ByChanderZap/api-basics/types"
	"github.com/ByChanderZap/api-basics/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Handler struct {
	store Querier
}

func NewHandler(store Querier) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.store.GetUserByEmail(r.Context(), payload.Email)
	if err == nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	} else if err != sql.ErrNoRows {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("something went wrong"))
		return
	}

	u, err := h.store.CreateUser(r.Context(), CreateUserParams{
		ID:        uuid.New(),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("unable to create new user"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, u)
}
