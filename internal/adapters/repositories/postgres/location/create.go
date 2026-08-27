package location

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *locationRepository) Create(ctx context.Context, loc *domain.Location) error {
	query := `
		INSERT INTO "flix-360".locations 
		(id, company_id, parent_id, name, code, type, status, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :parent_id, :name, :code, :type, :status, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, loc)
	return err
}
