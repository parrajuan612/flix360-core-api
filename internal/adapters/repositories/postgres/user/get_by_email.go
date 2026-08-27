package user

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	query := `SELECT * FROM "flix-360".users WHERE email = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &u, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("credenciales incorrectas") // Mensaje genérico por seguridad
		}
		return nil, err
	}
	return &u, nil
}
