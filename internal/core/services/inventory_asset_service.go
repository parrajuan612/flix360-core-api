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

type inventoryAssetService struct {
	repo ports.InventoryAssetRepository
}

func NewInventoryAssetService(repo ports.InventoryAssetRepository) ports.InventoryAssetService {
	return &inventoryAssetService{repo: repo}
}

func (s *inventoryAssetService) CreateAsset(ctx context.Context, asset *domain.InventoryAsset) error {
	if strings.TrimSpace(asset.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(asset.ProductID) == "" {
		return errors.New("el product_id es obligatorio")
	}
	if asset.Quantity <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	asset.ID = uuid.New().String()

	if asset.InventoryMode == "" {
		asset.InventoryMode = "unit"
	}
	if asset.Status == "" {
		asset.Status = "pending" // Estado inicial antes de la confirmación física
	}

	now := time.Now().UTC()
	asset.CreatedAt = now
	asset.UpdatedAt = now

	return s.repo.Create(ctx, asset)
}

func (s *inventoryAssetService) GetAssetByID(ctx context.Context, id string) (*domain.InventoryAsset, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
func (s *inventoryAssetService) ListAssets(ctx context.Context, companyID string, locationID string, limit, offset int) (*domain.PaginatedResponse, error) {
	if limit <= 0 || limit > 100 {
		limit = 50 // Por defecto traemos 50 activos
	}
	if offset < 0 {
		offset = 0
	}

	assets, total, err := s.repo.List(ctx, companyID, locationID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResponse{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Data:   assets,
	}, nil
}
