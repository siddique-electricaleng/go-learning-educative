package products

import (
	"context"
	repo "ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (repo.Product, error)
}

// Dependency injection here of repository
type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo *repo.Queries) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)

}

func (s *svc) GetProductByID(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.GetProductByID(ctx, id)

}
