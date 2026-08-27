package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/google/uuid"
)

type syncQueueService struct {
	repo ports.SyncQueueRepository
}

func NewSyncQueueService(repo ports.SyncQueueRepository) ports.SyncQueueService {
	return &syncQueueService{repo: repo}
}

func (s *syncQueueService) QueueJob(ctx context.Context, job *domain.SyncQueue) error {
	if strings.TrimSpace(job.CompanyID) == "" {
		return errors.New("company_id es obligatorio")
	}
	if strings.TrimSpace(job.EntityType) == "" || strings.TrimSpace(job.Operation) == "" {
		return errors.New("entity_type y operation son obligatorios")
	}
	if len(job.Payload) == 0 {
		return errors.New("el payload no puede estar vacío")
	}

	job.ID = uuid.New().String()
	job.Status = "pending" // Todo entra como pendiente para ser procesado después
	job.Attempts = 0

	now := time.Now().UTC()
	job.CreatedAt = now
	job.UpdatedAt = now

	return s.repo.Create(ctx, job)
}
