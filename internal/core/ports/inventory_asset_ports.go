package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type InventoryAssetRepository interface {
	Create(ctx context.Context, asset *domain.InventoryAsset) error
	GetByID(ctx context.Context, id string) (*domain.InventoryAsset, error)
	// Nueva función para listar:
	List(ctx context.Context, companyID string, locationID string, limit, offset int) ([]*domain.InventoryAsset, int64, error)
}

type InventoryAssetService interface {
	CreateAsset(ctx context.Context, asset *domain.InventoryAsset) error
	GetAssetByID(ctx context.Context, id string) (*domain.InventoryAsset, error)
	// Nueva función para listar:
	ListAssets(ctx context.Context, companyID string, locationID string, limit, offset int) (*domain.PaginatedResponse, error)
}
