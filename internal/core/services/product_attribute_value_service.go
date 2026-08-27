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

type productAttributeValueService struct {
	repo ports.ProductAttributeValueRepository
}

func NewProductAttributeValueService(repo ports.ProductAttributeValueRepository) ports.ProductAttributeValueService {
	return &productAttributeValueService{repo: repo}
}

func (s *productAttributeValueService) CreateValue(ctx context.Context, pav *domain.ProductAttributeValue) error {
	if strings.TrimSpace(pav.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(pav.ProductID) == "" {
		return errors.New("el product_id es obligatorio")
	}
	if strings.TrimSpace(pav.AttributeDefinitionID) == "" {
		return errors.New("el attribute_definition_id es obligatorio")
	}
	if len(pav.Value) == 0 {
		return errors.New("el valor (value) no puede estar vacío")
	}

	pav.ID = uuid.New().String()

	now := time.Now().UTC()
	pav.CreatedAt = now
	pav.UpdatedAt = now

	return s.repo.Create(ctx, pav)
}

func (s *productAttributeValueService) GetValueByID(ctx context.Context, id string) (*domain.ProductAttributeValue, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
