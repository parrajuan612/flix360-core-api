package attribute_definition

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *attributeDefinitionRepository) GetByID(ctx context.Context, id string) (*domain.AttributeDefinition, error) {
	var attr domain.AttributeDefinition
	query := `SELECT * FROM "flix-360".attribute_definitions WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &attr, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("definición de atributo no encontrada")
		}
		return nil, err
	}
	return &attr, nil
}
