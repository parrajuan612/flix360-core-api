package domain

import (
	"encoding/json"
	"time"
)

type SyncQueue struct {
	ID            string          `json:"id" db:"id"`
	CompanyID     string          `json:"company_id" db:"company_id"`
	DeviceID      *string         `json:"device_id,omitempty" db:"device_id"`
	EntityType    string          `json:"entity_type" db:"entity_type"` // ej: "inventory_movement"
	EntityID      string          `json:"entity_id" db:"entity_id"`
	Operation     string          `json:"operation" db:"operation"` // create, update, delete
	Payload       json.RawMessage `json:"payload" db:"payload"`     // El JSON completo que manda Android
	Status        string          `json:"status" db:"status"`       // pending, processing, done, failed
	Attempts      int             `json:"attempts" db:"attempts"`
	LastAttemptAt *time.Time      `json:"last_attempt_at,omitempty" db:"last_attempt_at"`
	ErrorMessage  *string         `json:"error_message,omitempty" db:"error_message"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
}
