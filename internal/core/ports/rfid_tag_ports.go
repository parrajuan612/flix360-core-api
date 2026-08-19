package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

// RfidTagRepository define las operaciones de base de datos
type RfidTagRepository interface {
	Create(ctx context.Context, tag *domain.RfidTag) error
	GetByEPC(ctx context.Context, epc string) (*domain.RfidTag, error)
}

// RfidTagService define los casos de uso del negocio
type RfidTagService interface {
	CreateTag(ctx context.Context, tag *domain.RfidTag) error
	GetTagByEPC(ctx context.Context, epc string) (*domain.RfidTag, error)
}
