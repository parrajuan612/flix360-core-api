package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

// ProductRepository es el puerto de salida (Outbound Port)
// Define las operaciones que la base de datos DEBE implementar.
type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	// Aquí a futuro agregaremos GetBySKU, List, Update, etc.
}

// ProductService es el puerto de entrada (Inbound Port)
// Define los casos de uso que los controladores HTTP pueden llamar.
type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
}
