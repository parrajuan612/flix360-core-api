package inventory_asset

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type inventoryAssetRepository struct {
	db *sqlx.DB
}

func NewInventoryAssetRepository(db *sqlx.DB) ports.InventoryAssetRepository {
	return &inventoryAssetRepository{
		db: db,
	}
}
