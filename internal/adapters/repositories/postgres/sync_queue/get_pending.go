package sync_queue

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *syncQueueRepository) GetPendingJobs(ctx context.Context, limit int) ([]*domain.SyncQueue, error) {
	var jobs []*domain.SyncQueue
	// Buscamos los trabajos pendientes, ordenados del más viejo al más nuevo
	query := `
		SELECT * FROM "flix-360".sync_queue 
		WHERE status = 'pending' 
		ORDER BY created_at ASC 
		LIMIT $1
	`
	err := r.db.SelectContext(ctx, &jobs, query, limit)
	return jobs, err
}
