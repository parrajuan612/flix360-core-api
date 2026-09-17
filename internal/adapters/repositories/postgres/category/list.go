package category

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *categoryRepository) List(ctx context.Context, companyID string, limit, offset int) ([]*domain.Category, int64, error) {
	var categories []*domain.Category
	var total int64

	err := r.db.GetContext(ctx, &total, `SELECT COUNT(*) FROM "flix-360".categories WHERE company_id = $1 AND status = 'active'`, companyID)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return []*domain.Category{}, 0, nil
	}

	query := `SELECT * FROM "flix-360".categories WHERE company_id = $1 AND status = 'active' ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	err = r.db.SelectContext(ctx, &categories, query, companyID, limit, offset)
	return categories, total, err
}
