package attribute_definition

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *attributeDefinitionRepository) Create(ctx context.Context, attr *domain.AttributeDefinition) error {
	query := `
		INSERT INTO "flix-360".attribute_definitions 
		(id, company_id, category_id, name, code, data_type, is_required, is_filterable, sort_order, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :category_id, :name, :code, :data_type, :is_required, :is_filterable, :sort_order, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, attr)
	return err
}
