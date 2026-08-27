package location

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) ports.LocationRepository {
	return &locationRepository{
		db: db,
	}
}
