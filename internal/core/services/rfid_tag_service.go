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

type rfidTagService struct {
	repo ports.RfidTagRepository
}

func NewRfidTagService(repo ports.RfidTagRepository) ports.RfidTagService {
	return &rfidTagService{repo: repo}
}

func (s *rfidTagService) CreateTag(ctx context.Context, tag *domain.RfidTag) error {
	if strings.TrimSpace(tag.EPC) == "" {
		return errors.New("el código EPC es obligatorio")
	}
	if strings.TrimSpace(tag.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}

	// Reglas de negocio para un tag "virgen"
	tag.ID = uuid.New().String()
	tag.Status = "unused" // Estado inicial

	now := time.Now().UTC()
	tag.CreatedAt = now
	tag.UpdatedAt = now

	return s.repo.Create(ctx, tag)
}

func (s *rfidTagService) GetTagByEPC(ctx context.Context, epc string) (*domain.RfidTag, error) {
	if strings.TrimSpace(epc) == "" {
		return nil, errors.New("el EPC no puede estar vacío")
	}
	return s.repo.GetByEPC(ctx, epc)
}
