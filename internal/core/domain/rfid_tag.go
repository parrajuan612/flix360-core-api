package domain

import "time"

type RfidTag struct {
	ID         string     `db:"id" json:"id"`
	CompanyID  string     `db:"company_id" json:"company_id"`
	EPC        string     `db:"epc" json:"epc"`
	Status     string     `db:"status" json:"status"` // unused, active, etc.
	AssignedAt *time.Time `db:"assigned_at" json:"assigned_at,omitempty"`
	RetiredAt  *time.Time `db:"retired_at" json:"retired_at,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
}
