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

type attributeDefinitionService struct {
	repo ports.AttributeDefinitionRepository
}

func NewAttributeDefinitionService(repo ports.AttributeDefinitionRepository) ports.AttributeDefinitionService {
	return &attributeDefinitionService{repo: repo}
}

func (s *attributeDefinitionService) CreateAttribute(ctx context.Context, attr *domain.AttributeDefinition) error {
	if strings.TrimSpace(attr.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(attr.Name) == "" {
		return errors.New("el nombre del atributo es obligatorio")
	}

	attr.ID = uuid.New().String()

	// Si no mandan tipo de dato, por defecto es texto
	if attr.DataType == "" {
		attr.DataType = "text"
	}

	now := time.Now().UTC()
	attr.CreatedAt = now
	attr.UpdatedAt = now

	return s.repo.Create(ctx, attr)
}

func (s *attributeDefinitionService) GetAttributeByID(ctx context.Context, id string) (*domain.AttributeDefinition, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
