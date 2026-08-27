package sync_queue

import (
	"context"
)

func (r *syncQueueRepository) UpdateJobStatus(ctx context.Context, id string, status string, errorMsg *string) error {
	query := `
		UPDATE "flix-360".sync_queue 
		SET status = $1, 
		    error_message = $2, 
		    attempts = attempts + 1, 
		    last_attempt_at = NOW(),
		    updated_at = NOW() 
		WHERE id = $3
	`
	_, err := r.db.ExecContext(ctx, query, status, errorMsg, id)
	return err
}
