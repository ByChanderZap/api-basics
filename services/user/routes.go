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
	/*
		TODO:
		I need to find a way to handle the errors on the unmarshall, i dont like to respond on the api with messages like:
		    "error": "json: cannot unmarshal number into Go struct field RegisterUserPayload.password of type string"

	*/
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, e := range errors {
			errorMessages[e.Field()] = e.Translate(utils.Translator)
		}
		utils.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":  "invalid payload",
			"fields": errorMessages,
		})
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
