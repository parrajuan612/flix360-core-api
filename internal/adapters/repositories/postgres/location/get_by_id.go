package location

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *locationRepository) GetByID(ctx context.Context, id string) (*domain.Location, error) {
	var loc domain.Location
	query := `SELECT * FROM "flix-360".locations WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &loc, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("locación no encontrada")
		}
		return nil, err
	}
	return &loc, nil
}
