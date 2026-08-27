package postgres

import (
	"context"
	"database/sql"
	"errors"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

// NewProductRepository es el constructor que inyecta la conexión a la base de datos
func NewProductRepository(db *sqlx.DB) ports.ProductRepository {
	return &productRepository{
		db: db,
	}
}

// Create inserta un nuevo producto en la base de datos
func (r *productRepository) Create(ctx context.Context, p *domain.Product) error {
	// Usamos el esquema "flix-360" tal como está en tu script SQL.
	// sqlx se encarga de reemplazar los valores con ":" por los campos de tu Struct gracias a las etiquetas `db`.
	query := `
		INSERT INTO "flix-360".products 
		(id, company_id, category_id, sku, name, description, inventory_mode, status, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :category_id, :sku, :name, :description, :inventory_mode, :status, :created_at, :updated_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, p)
	if err != nil {
		return err
	}

	return nil
}

// GetByID busca un producto por su ID
func (r *productRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product

	query := `
		SELECT 
			id, company_id, category_id, sku, name, description, inventory_mode, status, created_at, updated_at 
		FROM "flix-360".products 
		WHERE id = $1 
		LIMIT 1
	`

	// GetContext ejecuta la consulta y mapea el resultado directamente a la variable "product"
	err := r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		// Si no encuentra el registro, sqlx devuelve sql.ErrNoRows. Lo capturamos para dar un mensaje claro.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("producto no encontrado en la base de datos")
		}
		return nil, err
	}

	return &product, nil
}
func (r *productRepository) List(ctx context.Context, companyID string, limit, offset int, search string) ([]*domain.Product, int64, error) {
	var products []*domain.Product
	var total int64

	// 1. Contar el total de registros (útil para que el frontend dibuje los botones de las páginas)
	countQuery := `
		SELECT COUNT(*) 
		FROM "flix-360".products 
		WHERE company_id = $1 AND ($2 = '' OR name ILIKE '%' || $2 || '%' OR sku ILIKE '%' || $2 || '%')
	`
	err := r.db.GetContext(ctx, &total, countQuery, companyID, search)
	if err != nil {
		return nil, 0, err
	}

	// Si no hay ninguno, no hacemos la segunda consulta
	if total == 0 {
		return []*domain.Product{}, 0, nil
	}

	// 2. Traer los registros paginados
	query := `
		SELECT 
			id, company_id, category_id, sku, name, description, inventory_mode, status, created_at, updated_at 
		FROM "flix-360".products 
		WHERE company_id = $1 AND ($2 = '' OR name ILIKE '%' || $2 || '%' OR sku ILIKE '%' || $2 || '%')
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`

	err = r.db.SelectContext(ctx, &products, query, companyID, search, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
