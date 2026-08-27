package domain

import "time"

type Device struct {
	ID            string     `json:"id" db:"id"`
	CompanyID     string     `json:"company_id" db:"company_id"`
	Name          string     `json:"name" db:"name"`
	SerialNumber  string     `json:"serial_number" db:"serial_number"`
	Model         string     `json:"model,omitempty" db:"model"`
	Status        string     `json:"status" db:"status"` // active, inactive, blocked, maintenance, lost
	LastSyncAt    *time.Time `json:"last_sync_at,omitempty" db:"last_sync_at"`
	BlockedAt     *time.Time `json:"blocked_at,omitempty" db:"blocked_at"`
	BlockedReason string     `json:"blocked_reason,omitempty" db:"blocked_reason"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}
