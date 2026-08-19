package rfid_tag

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type rfidTagRepository struct {
	db *sqlx.DB
}

// NewRfidTagRepository instancia el repositorio
func NewRfidTagRepository(db *sqlx.DB) ports.RfidTagRepository {
	return &rfidTagRepository{
		db: db,
	}
}
