package attribute_definition

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type attributeDefinitionRepository struct {
	db *sqlx.DB
}

func NewAttributeDefinitionRepository(db *sqlx.DB) ports.AttributeDefinitionRepository {
	return &attributeDefinitionRepository{db: db}
}
