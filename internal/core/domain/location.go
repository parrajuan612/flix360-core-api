package domain

import "time"

type Location struct {
	ID        string    `db:"id" json:"id"`
	CompanyID string    `db:"company_id" json:"company_id"`
	ParentID  *string   `db:"parent_id" json:"parent_id,omitempty"`
	Name      string    `db:"name" json:"name"`
	Code      *string   `db:"code" json:"code,omitempty"`
	Type      string    `db:"type" json:"type"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
