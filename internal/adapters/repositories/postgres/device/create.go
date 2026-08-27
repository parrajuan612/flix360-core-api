package device

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *deviceRepository) Create(ctx context.Context, dev *domain.Device) error {
	query := `
		INSERT INTO "flix-360".devices 
		(id, company_id, name, serial_number, model, status, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :name, :serial_number, :model, :status, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, dev)
	return err
}
