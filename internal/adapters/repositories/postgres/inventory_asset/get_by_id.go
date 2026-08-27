package inventory_asset

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *inventoryAssetRepository) GetByID(ctx context.Context, id string) (*domain.InventoryAsset, error) {
	var asset domain.InventoryAsset
	query := `SELECT * FROM "flix-360".inventory_assets WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &asset, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("activo de inventario no encontrado")
		}
		return nil, err
	}
	return &asset, nil
}
