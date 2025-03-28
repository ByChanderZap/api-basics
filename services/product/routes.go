package product

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
	router.Put("/products/{id}", h.handleUpdateProduct)
	router.Delete("/products/{id}", h.handleDeleteProduct)
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
		Quantity:    int32(payload.Quantity),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Println("Error creating Product", err)
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("unable to create product"))
		return
	}

	response := types.ProductResponse{
		ID:          p.ID.String(),
		Name:        p.Name,
		Description: p.Description,
		Image:       utils.NullableString(p.Image),
		Price:       p.Price,
		Quantity:    p.Quantity,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   utils.NullableTime(p.DeletedAt),
	}

	utils.WriteJSON(w, http.StatusCreated, response)
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

func (h *Handler) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	pId := chi.URLParam(r, "id")
	parsedId, err := uuid.Parse(pId)

	if err != nil {
		log.Println("Error parsing product id", err)
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid product id"))
		return
	}

	var payload types.UpdateProductPayload
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

	updated, err := h.store.UpdateProduct(r.Context(), UpdateProductParams{
		ID:          parsedId,
		Name:        payload.Name,
		Description: payload.Description,
		Image:       img,
		Price:       payload.Price,
		Quantity:    int32(payload.Quantity),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Println("Error updating product", err)
		if strings.Contains(err.Error(), "no rows in result set") {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Errorf("Product with id %s not found", pId))
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, errors.New("unable to update product"))
		return
	}

	toResponse := types.ProductResponse{
		ID:          updated.ID.String(),
		Name:        updated.Name,
		Description: updated.Description,
		Image:       utils.NullableString(updated.Image),
		Price:       updated.Price,
		Quantity:    updated.Quantity,
		CreatedAt:   updated.CreatedAt,
		UpdatedAt:   updated.UpdatedAt,
		DeletedAt:   utils.NullableTime(updated.DeletedAt),
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "product updated",
		"product": toResponse,
	})
}

func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	pId := chi.URLParam(r, "id")
	parsedId, err := uuid.Parse(pId)

	if err != nil {
		log.Println("Error parsing product id", err)
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid product id"))
		return
	}

	_, err = h.store.DeleteProduct(r.Context(), DeleteProductParams{
		ID:        parsedId,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Error deleting product", err)
		if strings.Contains(err.Error(), "no rows in result set") {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Errorf("Product with id %s not found", pId))
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, errors.New("unable to delete product"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "product with id " + pId + " deleted",
	})
}
