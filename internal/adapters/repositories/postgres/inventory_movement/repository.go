package inventory_movement

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type inventoryMovementRepository struct {
	db *sqlx.DB
}

func NewInventoryMovementRepository(db *sqlx.DB) ports.InventoryMovementRepository {
	return &inventoryMovementRepository{
		db: db,
	}
}
