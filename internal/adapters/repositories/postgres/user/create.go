package user

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	query := `
		INSERT INTO "flix-360".users 
		(id, company_id, full_name, email, password_hash, role, status, created_at, updated_at) 
		VALUES 
		(:id, :company_id, :full_name, :email, :password_hash, :role, :status, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, u)
	return err
}
