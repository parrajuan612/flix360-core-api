package device

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *deviceRepository) GetByID(ctx context.Context, id string) (*domain.Device, error) {
	var dev domain.Device
	query := `SELECT * FROM "flix-360".devices WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &dev, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("dispositivo no encontrado")
		}
		return nil, err
	}
	return &dev, nil
}
