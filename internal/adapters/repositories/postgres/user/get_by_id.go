package user

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var u domain.User
	query := `SELECT * FROM "flix-360".users WHERE id = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &u, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}
	return &u, nil
}
