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

type inventoryMovementService struct {
	repo ports.InventoryMovementRepository
}

func NewInventoryMovementService(repo ports.InventoryMovementRepository) ports.InventoryMovementService {
	return &inventoryMovementService{repo: repo}
}

func (s *inventoryMovementService) CreateMovement(ctx context.Context, mov *domain.InventoryMovement) error {
	if strings.TrimSpace(mov.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(mov.AssetID) == "" {
		return errors.New("el asset_id es obligatorio")
	}
	if strings.TrimSpace(mov.MovementType) == "" {
		return errors.New("el tipo de movimiento (movement_type) es obligatorio")
	}
	if mov.Quantity <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	// Autogeneramos IDs y fechas
	mov.ID = uuid.New().String()

	now := time.Now().UTC()
	if mov.PerformedAt.IsZero() {
		mov.PerformedAt = now
	}
	mov.CreatedAt = now
	mov.UpdatedAt = now

	return s.repo.Create(ctx, mov)
}

func (s *inventoryMovementService) GetMovementByID(ctx context.Context, id string) (*domain.InventoryMovement, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
func (s *inventoryMovementService) CreateBulk(ctx context.Context, companyID, userID string, movs []*domain.InventoryMovement) error {
	if len(movs) == 0 {
		return errors.New("la lista de movimientos está vacía")
	}

	now := time.Now().UTC()

	// Recorremos el arreglo de 100 o 500 lecturas
	for _, mov := range movs {
		// Validaciones básicas por cada registro
		if strings.TrimSpace(mov.AssetID) == "" {
			return errors.New("todos los movimientos deben tener un asset_id")
		}

		// Autogeneramos IDs y Fechas
		mov.ID = uuid.New().String()

		// ¡SEGURIDAD! Forzamos el company_id y user_id desde el Token (ignora lo que mande Android)
		mov.CompanyID = companyID
		mov.PerformedBy = &userID // Apuntamos al ID del usuario que hizo login

		if mov.PerformedAt.IsZero() {
			mov.PerformedAt = now
		}
		mov.CreatedAt = now
		mov.UpdatedAt = now
	}

	return s.repo.CreateBulk(ctx, movs)
}
