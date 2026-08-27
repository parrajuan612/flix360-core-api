package product_attribute_value

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type productAttributeValueRepository struct {
	db *sqlx.DB
}

func NewProductAttributeValueRepository(db *sqlx.DB) ports.ProductAttributeValueRepository {
	return &productAttributeValueRepository{db: db}
}
