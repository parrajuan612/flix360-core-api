package domain

import "time"

type InventoryMovement struct {
	ID             string    `db:"id" json:"id"`
	CompanyID      string    `db:"company_id" json:"company_id"`
	AssetID        string    `db:"asset_id" json:"asset_id"`
	DeviceID       *string   `db:"device_id" json:"device_id,omitempty"`
	FromLocationID *string   `db:"from_location_id" json:"from_location_id,omitempty"`
	ToLocationID   *string   `db:"to_location_id" json:"to_location_id,omitempty"`
	MovementType   string    `db:"movement_type" json:"movement_type"`
	Quantity       float64   `db:"quantity" json:"quantity"`
	ReferenceType  *string   `db:"reference_type" json:"reference_type,omitempty"`
	ReferenceID    *string   `db:"reference_id" json:"reference_id,omitempty"`
	Notes          *string   `db:"notes" json:"notes,omitempty"`
	PerformedBy    *string   `db:"performed_by" json:"performed_by,omitempty"`
	PerformedAt    time.Time `db:"performed_at" json:"performed_at"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
