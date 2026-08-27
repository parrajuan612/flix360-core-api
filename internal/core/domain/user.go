package domain

import "time"

type User struct {
	ID            string     `json:"id" db:"id"`
	CompanyID     string     `json:"company_id" db:"company_id"`
	FullName      string     `json:"full_name" db:"full_name"`
	Email         string     `json:"email" db:"email"`
	PasswordHash  string     `json:"password_hash,omitempty" db:"password_hash"` // omitempty para no devolverlo en el JSON
	Role          string     `json:"role" db:"role"`                             // admin, manager, operator
	Status        string     `json:"status" db:"status"`                         // active, inactive, blocked
	LastLoginAt   *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	BlockedAt     *time.Time `json:"blocked_at,omitempty" db:"blocked_at"`
	BlockedReason *string    `json:"blocked_reason,omitempty" db:"blocked_reason"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}
