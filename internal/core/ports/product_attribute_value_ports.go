package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type ProductAttributeValueRepository interface {
	Create(ctx context.Context, pav *domain.ProductAttributeValue) error
	GetByID(ctx context.Context, id string) (*domain.ProductAttributeValue, error)
}

type ProductAttributeValueService interface {
	CreateValue(ctx context.Context, pav *domain.ProductAttributeValue) error
	GetValueByID(ctx context.Context, id string) (*domain.ProductAttributeValue, error)
}
