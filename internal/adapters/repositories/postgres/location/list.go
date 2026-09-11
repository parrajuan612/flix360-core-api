package location

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *locationRepository) List(ctx context.Context, companyID string, limit, offset int) ([]*domain.Location, int64, error) {
	var locations []*domain.Location
	var total int64

	// Contar totales
	countQuery := `SELECT COUNT(*) FROM "flix-360".locations WHERE company_id = $1 AND status = 'active'`
	err := r.db.GetContext(ctx, &total, countQuery, companyID)
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return []*domain.Location{}, 0, nil
	}

	// Traer registros
	query := `
		SELECT * FROM "flix-360".locations 
		WHERE company_id = $1 AND status = 'active'
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	err = r.db.SelectContext(ctx, &locations, query, companyID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return locations, total, nil
}
