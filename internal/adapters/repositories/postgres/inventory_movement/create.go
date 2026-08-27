package inventory_movement

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *inventoryMovementRepository) Create(ctx context.Context, mov *domain.InventoryMovement) error {
	query := `
		INSERT INTO "flix-360".inventory_movements 
		(id, company_id, asset_id, device_id, from_location_id, to_location_id, movement_type, quantity, reference_type, reference_id, notes, performed_by, performed_at, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :asset_id, :device_id, :from_location_id, :to_location_id, :movement_type, :quantity, :reference_type, :reference_id, :notes, :performed_by, :performed_at, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, mov)
	return err
}
