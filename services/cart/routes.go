package cart

import (
	"log"
	"net/http"

	"github.com/ByChanderZap/api-basics/services/product"
	"github.com/ByChanderZap/api-basics/types"
	"github.com/ByChanderZap/api-basics/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db           *pgxpool.Pool
	orderStore   *Queries
	productStore *product.Queries
}

func NewHandler(db *pgxpool.Pool, store *Queries, productStore *product.Queries) *Handler {
	return &Handler{
		db:           db,
		orderStore:   store,
		productStore: productStore,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Post("/cart/checkout", h.handleCheckout)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var reqCart types.CartCheckoutPayload

	if err := utils.ParseJson(r, &reqCart); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(reqCart); err != nil {
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

	productIds, err := getCartItemsIds(reqCart.Items)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	orderId, totalPrice, err := h.createOrder(productIds, reqCart.Items, uuid.UUID{})
	if err != nil {
		log.Println("error creating order", err)
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message":     "order created",
		"order_id":    orderId,
		"total_price": totalPrice,
	})
}
