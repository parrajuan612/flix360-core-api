package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type InventoryMovementRepository interface {
	Create(ctx context.Context, movement *domain.InventoryMovement) error
	GetByID(ctx context.Context, id string) (*domain.InventoryMovement, error)
	CreateBulk(ctx context.Context, movements []*domain.InventoryMovement) error // <-- Nueva
}

type InventoryMovementService interface {
	CreateMovement(ctx context.Context, movement *domain.InventoryMovement) error
	GetMovementByID(ctx context.Context, id string) (*domain.InventoryMovement, error)
	CreateBulk(ctx context.Context, companyID, userID string, movements []*domain.InventoryMovement) error // <-- Nueva
}
