package domain

import "time"

type AttributeDefinition struct {
	ID           string    `json:"id" db:"id"`
	CompanyID    string    `json:"company_id" db:"company_id"`
	CategoryID   *string   `json:"category_id,omitempty" db:"category_id"` // Puede ser null si es un atributo global
	Name         string    `json:"name" db:"name"`
	Code         *string   `json:"code,omitempty" db:"code"`
	DataType     string    `json:"data_type" db:"data_type"` // text, number, boolean, date, select, json
	IsRequired   bool      `json:"is_required" db:"is_required"`
	IsFilterable bool      `json:"is_filterable" db:"is_filterable"`
	SortOrder    int       `json:"sort_order" db:"sort_order"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
