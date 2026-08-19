package domain

import "time"

type Product struct {
	ID            string    `db:"id" json:"id"`
	CompanyID     string    `db:"company_id" json:"company_id"`
	CategoryID    *string   `db:"category_id" json:"category_id,omitempty"` // Es NULLable en tu SQL
	SKU           string    `db:"sku" json:"sku"`
	Name          string    `db:"name" json:"name"`
	Description   *string   `db:"description" json:"description,omitempty"`
	InventoryMode string    `db:"inventory_mode" json:"inventory_mode"` // Enum por defecto 'unit'
	Status        string    `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
