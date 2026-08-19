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

type productService struct {
	repo ports.ProductRepository
}

// NewProductService inyecta el repositorio a través de su interfaz
func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, product *domain.Product) error {
	// 1. Validaciones de negocio puras
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("el nombre del producto es obligatorio")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("el SKU del producto es obligatorio")
	}
	if strings.TrimSpace(product.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}

	// 2. Asignación de valores por defecto (evitamos delegar esto solo a la BD)
	product.ID = uuid.New().String()

	if product.InventoryMode == "" {
		product.InventoryMode = "unit" // Valor por defecto según tu SQL
	}
	if product.Status == "" {
		product.Status = "active" // Valor por defecto según tu SQL
	}

	now := time.Now().UTC()
	product.CreatedAt = now
	product.UpdatedAt = now

	// 3. Llamada al puerto de salida (Repositorio)
	return s.repo.Create(ctx, product)
}

func (s *productService) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	if id == "" {
		return nil, errors.New("el ID del producto no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
