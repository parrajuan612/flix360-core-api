package product_attribute_value

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *productAttributeValueRepository) GetByID(ctx context.Context, id string) (*domain.ProductAttributeValue, error) {
	var pav domain.ProductAttributeValue
	query := `SELECT * FROM "flix-360".product_attribute_values WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &pav, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("valor de atributo no encontrado")
		}
		return nil, err
	}
	return &pav, nil
}
