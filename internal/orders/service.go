package orders

import (
	"context"
	repo "ecom/internal/adapters/postgresql/sqlc"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product does not have enough stock")
)

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}

type svc struct {
	// repository
	repo *repo.Queries

	db *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	// validate payload
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer ID is required")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least one item is required")
	}

	// Transactions - taken from sqlc docs on transactions
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	// Rollingback the transaction
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	// create an order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}
	// look for the product if exists
	for _, item := range tempOrder.Items {
		product, err := qtx.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNoStock
		}
		// create order item
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:      order.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PriceInCents: product.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, err
		}

		// Update the Product Stock Quantity in the database
		productUpdatedQuantity := product.Quantity - item.Quantity

		err = qtx.UpdateProductQuantity(ctx, repo.UpdateProductQuantityParams{
			Quantity: productUpdatedQuantity,
			ID:       item.ProductID,
		})
		if err != nil {
			return repo.Order{}, err
		}
	}
	// Commit the transaction
	tx.Commit(ctx)

	return order, nil
}
