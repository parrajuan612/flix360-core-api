package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type AttributeDefinitionRepository interface {
	Create(ctx context.Context, attr *domain.AttributeDefinition) error
	GetByID(ctx context.Context, id string) (*domain.AttributeDefinition, error)
}

type AttributeDefinitionService interface {
	CreateAttribute(ctx context.Context, attr *domain.AttributeDefinition) error
	GetAttributeByID(ctx context.Context, id string) (*domain.AttributeDefinition, error)
}
