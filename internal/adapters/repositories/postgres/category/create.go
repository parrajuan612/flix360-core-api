package category

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *categoryRepository) Create(ctx context.Context, cat *domain.Category) error {
	query := `
		INSERT INTO "flix-360".categories 
		(id, company_id, parent_id, name, code, description, status, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :parent_id, :name, :code, :description, :status, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, cat)
	return err
}
