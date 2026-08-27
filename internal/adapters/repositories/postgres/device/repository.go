package device

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type deviceRepository struct {
	db *sqlx.DB
}

func NewDeviceRepository(db *sqlx.DB) ports.DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}
