package category

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *categoryRepository) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	var cat domain.Category
	query := `SELECT * FROM "flix-360".categories WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &cat, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("categoría no encontrada")
		}
		return nil, err
	}
	return &cat, nil
}
