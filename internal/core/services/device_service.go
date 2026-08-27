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

type deviceService struct {
	repo ports.DeviceRepository
}

func NewDeviceService(repo ports.DeviceRepository) ports.DeviceService {
	return &deviceService{repo: repo}
}

func (s *deviceService) CreateDevice(ctx context.Context, dev *domain.Device) error {
	if strings.TrimSpace(dev.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(dev.Name) == "" {
		return errors.New("el nombre del dispositivo es obligatorio")
	}
	if strings.TrimSpace(dev.SerialNumber) == "" {
		return errors.New("el número de serie es obligatorio")
	}

	dev.ID = uuid.New().String()

	if dev.Status == "" {
		dev.Status = "active"
	}

	now := time.Now().UTC()
	dev.CreatedAt = now
	dev.UpdatedAt = now

	return s.repo.Create(ctx, dev)
}

func (s *deviceService) GetDeviceByID(ctx context.Context, id string) (*domain.Device, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
