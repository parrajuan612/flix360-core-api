package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	GetByID(ctx context.Context, id string) (*domain.Category, error)
}

type CategoryService interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetCategoryByID(ctx context.Context, id string) (*domain.Category, error)
}
