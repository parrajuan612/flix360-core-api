package sync_queue

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *syncQueueRepository) Create(ctx context.Context, job *domain.SyncQueue) error {
	query := `
		INSERT INTO "flix-360".sync_queue 
		(id, company_id, device_id, entity_type, entity_id, operation, payload, status, attempts, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :device_id, :entity_type, :entity_id, :operation, :payload, :status, :attempts, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, job)
	return err
}
