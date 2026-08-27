package inventory_asset

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *inventoryAssetRepository) Create(ctx context.Context, asset *domain.InventoryAsset) error {
	query := `
		INSERT INTO "flix-360".inventory_assets 
		(id, company_id, product_id, rfid_tag_id, location_id, quantity, inventory_mode, status, batch_code, notes, created_by, updated_by, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :product_id, :rfid_tag_id, :location_id, :quantity, :inventory_mode, :status, :batch_code, :notes, :created_by, :updated_by, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, asset)
	return err
}
