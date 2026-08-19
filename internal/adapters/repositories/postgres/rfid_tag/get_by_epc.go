package rfid_tag

import (
	"context"
	"database/sql"
	"errors"
	"flix360-core-api/internal/core/domain"
)

func (r *rfidTagRepository) GetByEPC(ctx context.Context, epc string) (*domain.RfidTag, error) {
	var tag domain.RfidTag

	query := `SELECT * FROM "flix-360".rfid_tags WHERE epc = $1 LIMIT 1`

	err := r.db.GetContext(ctx, &tag, query, epc)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("etiqueta rfid no encontrada")
		}
		return nil, err
	}

	return &tag, nil
}
