package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type LocationRepository interface {
	Create(ctx context.Context, location *domain.Location) error
	GetByID(ctx context.Context, id string) (*domain.Location, error)
}

type LocationService interface {
	CreateLocation(ctx context.Context, location *domain.Location) error
	GetLocationByID(ctx context.Context, id string) (*domain.Location, error)
}
