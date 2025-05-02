package cart

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	cartStore "github.com/ByChanderZap/api-basics/services/cart/generated"
	productStore "github.com/ByChanderZap/api-basics/services/product/generated"
	"github.com/ByChanderZap/api-basics/types"
	"github.com/google/uuid"
)

func getCartItemsIds(items []types.CartItem) ([]uuid.UUID, error) {
	productIds := make([]uuid.UUID, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %v", item.ProductId)
		}

		productIds[i] = item.ProductId
	}
	return productIds, nil
}

func (h *Handler) createOrder(
	productsIds []uuid.UUID,
	cartItems []types.CartItem,
	userId uuid.UUID,
) (uuid.UUID, float64, error) {
	tx, err := h.db.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return uuid.UUID{}, 0, fmt.Errorf("something went wrong while starting the transaction")
	}
	defer tx.Rollback(context.Background())

	orderTx := h.orderStore.WithTx(tx)
	prodTx := h.productStore.WithTx(tx)

	ps, err := prodTx.GetProductsByIds(context.Background(), productsIds)
	if err != nil {
		log.Println("error getting products by ids", err)
		return uuid.UUID{}, 0, fmt.Errorf("error getting products by ids")
	}

	prodMap := make(map[uuid.UUID]productStore.GetProductsByIdsRow)
	for _, p := range ps {
		prodMap[p.ID] = p
	}

	if err := checkIfCartIsInStock(cartItems, prodMap); err != nil {
		return uuid.UUID{}, 0, err
	}

	totalPrice := calculateTotalCartPrice(cartItems, prodMap)

	orderId := uuid.New()

	order, err := orderTx.CreateOrder(context.Background(), cartStore.CreateOrderParams{
		ID:        orderId,
		UserID:    userId, //PLACE HOLDER
		Total:     totalPrice,
		Status:    cartStore.OrderStatusPending,
		Address:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return uuid.UUID{}, 0, err
	}

	for _, item := range cartItems {
		p := prodMap[item.ProductId]

		_, err := orderTx.CreateOrderItem(context.Background(), cartStore.CreateOrderItemParams{
			ID:        uuid.UUID{},
			OrderID:   order.ID,
			ProductID: p.ID,
			Quantity:  p.Quantity,
			Price:     p.Price * float64(p.Quantity),
		})
		if err != nil {
			return uuid.UUID{}, 0, err
		}

		_, err = prodTx.DecreaseProductStock(context.Background(), productStore.DecreaseProductStockParams{
			ID:        p.ID,
			Quantity:  p.Quantity,
			UpdatedAt: time.Now(),
		})
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.UUID{}, 0, fmt.Errorf("product %v is out of stock or already updated", p.ID)
		}
		if err != nil {
			return uuid.UUID{}, 0, fmt.Errorf("failed to decrease stock for product %v: %w", p.ID, err)
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		log.Println("something went wrong while committing the transaction", err)
		return uuid.UUID{}, 0, fmt.Errorf("something went wrong while processing the order")
	}

	return order.ID, order.Total, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[uuid.UUID]productStore.GetProductsByIdsRow) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		prod, ok := products[item.ProductId]
		if !ok {
			return fmt.Errorf("product %v not found", item.ProductId)
		}

		if prod.Quantity < int32(item.Quantity) {
			return fmt.Errorf("product %v is out of stock", item.ProductId)
		}
	}
	return nil
}

func calculateTotalCartPrice(cartItems []types.CartItem, products map[uuid.UUID]productStore.GetProductsByIdsRow) float64 {
	var total float64

	for _, item := range cartItems {
		prod := products[item.ProductId]
		total += prod.Price * float64(prod.Quantity)
	}

	return total
}
