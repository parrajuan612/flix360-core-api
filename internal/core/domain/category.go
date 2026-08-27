package domain

import "time"

type Category struct {
	ID          string    `json:"id" db:"id"`
	CompanyID   string    `json:"company_id" db:"company_id"`
	ParentID    *string   `json:"parent_id,omitempty" db:"parent_id"` // Para subcategorías (ej. Ropa -> Deportiva)
	Name        string    `json:"name" db:"name"`
	Code        *string   `json:"code,omitempty" db:"code"`
	Description *string   `json:"description,omitempty" db:"description"`
	Status      string    `json:"status" db:"status"` // active, inactive
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
