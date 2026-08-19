package domain

import "time"

type InventoryAsset struct {
	ID            string    `db:"id" json:"id"`
	CompanyID     string    `db:"company_id" json:"company_id"`
	ProductID     string    `db:"product_id" json:"product_id"`
	RfidTagID     *string   `db:"rfid_tag_id" json:"rfid_tag_id,omitempty"`
	LocationID    *string   `db:"location_id" json:"location_id,omitempty"`
	Quantity      float64   `db:"quantity" json:"quantity"`
	InventoryMode string    `db:"inventory_mode" json:"inventory_mode"`
	Status        string    `db:"status" json:"status"`
	BatchCode     *string   `db:"batch_code" json:"batch_code,omitempty"`
	Notes         *string   `db:"notes" json:"notes,omitempty"`
	CreatedBy     *string   `db:"created_by" json:"created_by,omitempty"`
	UpdatedBy     *string   `db:"updated_by" json:"updated_by,omitempty"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
