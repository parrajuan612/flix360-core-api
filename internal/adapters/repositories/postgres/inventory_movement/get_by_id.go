package inventory_movement

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *inventoryMovementRepository) GetByID(ctx context.Context, id string) (*domain.InventoryMovement, error) {
	var mov domain.InventoryMovement
	query := `SELECT * FROM "flix-360".inventory_movements WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &mov, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("movimiento de inventario no encontrado")
		}
		return nil, err
	}
	return &mov, nil
}
