package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type DeviceRepository interface {
	Create(ctx context.Context, device *domain.Device) error
	GetByID(ctx context.Context, id string) (*domain.Device, error)
}

type DeviceService interface {
	CreateDevice(ctx context.Context, device *domain.Device) error
	GetDeviceByID(ctx context.Context, id string) (*domain.Device, error)
}
