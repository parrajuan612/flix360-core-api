package product_attribute_value

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *productAttributeValueRepository) Create(ctx context.Context, pav *domain.ProductAttributeValue) error {
	query := `
		INSERT INTO "flix-360".product_attribute_values 
		(id, company_id, product_id, attribute_definition_id, value, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :product_id, :attribute_definition_id, :value, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, pav)
	return err
}
