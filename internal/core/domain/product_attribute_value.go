package domain

import (
	"encoding/json"
	"time"
)

type ProductAttributeValue struct {
	ID                    string          `json:"id" db:"id"`
	CompanyID             string          `json:"company_id" db:"company_id"`
	ProductID             string          `json:"product_id" db:"product_id"`
	AttributeDefinitionID string          `json:"attribute_definition_id" db:"attribute_definition_id"`
	Value                 json.RawMessage `json:"value" db:"value"`
	CreatedAt             time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at" db:"updated_at"`
}
