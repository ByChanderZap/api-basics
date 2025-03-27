package product

import (
	"database/sql"
	"log"
	"net/http"

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
	router.Get("/products", h.handleGetProducts)
	router.Post("/products", h.handleCreateProduct)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload
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
	img := sql.NullString{}
	if payload.Image != "" {
		img.Valid = true
		img.String = payload.Image
	}
	p, err := h.store.CreateProduct(r.Context(), CreateProductParams{
		ID:          uuid.New(),
		Name:        payload.Name,
		Description: payload.Description,
		Image:       img,
		Price:       payload.Price,
	})
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	p, err := h.store.GetProducts(r.Context())
	if err != nil {
		log.Println("Error getting products", err)
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	if p == nil {
		utils.WriteJSON(w, http.StatusOK, []interface{}{})
		return
	}
	utils.WriteJSON(w, http.StatusOK, p)
}
