package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"
)

type SyncWorker struct {
	syncRepo ports.SyncQueueRepository
	movSvc   ports.InventoryMovementService
}

func NewSyncWorker(syncRepo ports.SyncQueueRepository, movSvc ports.InventoryMovementService) *SyncWorker {
	return &SyncWorker{
		syncRepo: syncRepo,
		movSvc:   movSvc,
	}
}

// Start inicia el ciclo infinito del worker
func (w *SyncWorker) Start(ctx context.Context) {
	log.Println("🚀 Sincronizador Offline iniciado en segundo plano...")

	// El "Ticker" hace que el ciclo se ejecute cada 10 segundos
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // Si apagamos el servidor, el worker se detiene
			log.Println("🛑 Sincronizador Offline detenido.")
			return
		case <-ticker.C:
			w.processPendingJobs(ctx)
		}
	}
}

func (w *SyncWorker) processPendingJobs(ctx context.Context) {
	// Traemos hasta 50 trabajos pendientes
	jobs, err := w.syncRepo.GetPendingJobs(ctx, 50)
	if err != nil {
		log.Printf("Error buscando trabajos pendientes: %v", err)
		return
	}

	if len(jobs) == 0 {
		return // No hay nada que hacer, volvemos a dormir
	}

	log.Printf("⚙️ Procesando %d trabajos offline pendientes...", len(jobs))

	for _, job := range jobs {
		var processErr error

		// Dependiendo de qué entidad nos mandó Android, lo procesamos
		switch job.EntityType {
		case "inventory_movement":
			processErr = w.processInventoryMovement(ctx, job)
		default:
			processErr = fmt.Errorf("tipo de entidad no soportado: %s", job.EntityType)
		}

		// Si hubo un error, marcamos como fallido, sino como terminado
		if processErr != nil {
			errMsg := processErr.Error()
			w.syncRepo.UpdateJobStatus(ctx, job.ID, "failed", &errMsg)
			log.Printf("❌ Error procesando trabajo %s: %s", job.ID, errMsg)
		} else {
			w.syncRepo.UpdateJobStatus(ctx, job.ID, "done", nil)
			log.Printf("✅ Trabajo %s procesado con éxito", job.ID)
		}
	}
}

func (w *SyncWorker) processInventoryMovement(ctx context.Context, job *domain.SyncQueue) error {
	var mov domain.InventoryMovement
	// Convertimos el JSON (Payload) a nuestro struct de Movimiento
	if err := json.Unmarshal(job.Payload, &mov); err != nil {
		return fmt.Errorf("JSON malformado en payload: %v", err)
	}

	// Como es offline, debemos asegurarnos de que la empresa se asigne correctamente
	mov.CompanyID = job.CompanyID

	// Guardamos el movimiento usando el servicio normal
	return w.movSvc.CreateMovement(ctx, &mov)
}
