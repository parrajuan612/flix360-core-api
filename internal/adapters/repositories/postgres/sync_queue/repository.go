package sync_queue

import (
	"flix360-core-api/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type syncQueueRepository struct {
	db *sqlx.DB
}

func NewSyncQueueRepository(db *sqlx.DB) ports.SyncQueueRepository {
	return &syncQueueRepository{db: db}
}
