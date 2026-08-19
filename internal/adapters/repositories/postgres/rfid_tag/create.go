package rfid_tag

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *rfidTagRepository) Create(ctx context.Context, tag *domain.RfidTag) error {
	query := `
		INSERT INTO "flix-360".rfid_tags 
		(id, company_id, epc, status, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :epc, :status, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, tag)
	return err
}
