package inventory_asset

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *inventoryAssetRepository) List(ctx context.Context, companyID string, locationID string, limit, offset int) ([]*domain.InventoryAsset, int64, error) {
	var assets []*domain.InventoryAsset
	var total int64

	// 1. Construimos el contador dinámicamente
	countQuery := `SELECT COUNT(*) FROM "flix-360".inventory_assets WHERE company_id = $1`

	var err error
	if locationID != "" {
		countQuery += ` AND location_id = $2`
		err = r.db.GetContext(ctx, &total, countQuery, companyID, locationID)
	} else {
		err = r.db.GetContext(ctx, &total, countQuery, companyID)
	}

	if err != nil {
		// Loguear el error real sería ideal aquí si tuviéramos un sistema de logs
		return nil, 0, err
	}

	if total == 0 {
		return []*domain.InventoryAsset{}, 0, nil
	}

	// 2. Construimos el listado dinámicamente
	query := `SELECT * FROM "flix-360".inventory_assets WHERE company_id = $1`

	if locationID != "" {
		query += ` AND location_id = $2 ORDER BY created_at DESC LIMIT $3 OFFSET $4`
		err = r.db.SelectContext(ctx, &assets, query, companyID, locationID, limit, offset)
	} else {
		query += ` ORDER BY created_at DESC LIMIT $2 OFFSET $3`
		err = r.db.SelectContext(ctx, &assets, query, companyID, limit, offset)
	}

	if err != nil {
		return nil, 0, err
	}

	return assets, total, nil
}
