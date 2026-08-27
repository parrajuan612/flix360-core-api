package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	// Nueva función para listar:
	List(ctx context.Context, companyID string, limit, offset int, search string) ([]*domain.Product, int64, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	// Nueva función para listar:
	ListProducts(ctx context.Context, companyID string, limit, offset int, search string) (*domain.PaginatedResponse, error)
}
