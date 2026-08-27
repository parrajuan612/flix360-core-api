package category

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) ports.CategoryRepository {
	return &categoryRepository{db: db}
}
