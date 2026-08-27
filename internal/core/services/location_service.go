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

type locationService struct {
	repo ports.LocationRepository
}

func NewLocationService(repo ports.LocationRepository) ports.LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) CreateLocation(ctx context.Context, loc *domain.Location) error {
	if strings.TrimSpace(loc.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(loc.Name) == "" {
		return errors.New("el nombre de la locación es obligatorio")
	}

	loc.ID = uuid.New().String()

	// Valores por defecto de tu base de datos
	if loc.Type == "" {
		loc.Type = "store" // Puede ser warehouse, zone, store, etc.
	}
	if loc.Status == "" {
		loc.Status = "active"
	}

	now := time.Now().UTC()
	loc.CreatedAt = now
	loc.UpdatedAt = now

	return s.repo.Create(ctx, loc)
}

func (s *locationService) GetLocationByID(ctx context.Context, id string) (*domain.Location, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
