package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type SyncQueueRepository interface {
	Create(ctx context.Context, job *domain.SyncQueue) error
	GetPendingJobs(ctx context.Context, limit int) ([]*domain.SyncQueue, error)            // <-- Nueva
	UpdateJobStatus(ctx context.Context, id string, status string, errorMsg *string) error // <-- Nueva
}

type SyncQueueService interface {
	QueueJob(ctx context.Context, job *domain.SyncQueue) error
}
