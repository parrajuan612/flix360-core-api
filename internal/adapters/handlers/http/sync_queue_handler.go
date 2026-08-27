package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type SyncQueueHandler struct {
	service ports.SyncQueueService
}

func NewSyncQueueHandler(service ports.SyncQueueService) *SyncQueueHandler {
	return &SyncQueueHandler{service: service}
}

func (h *SyncQueueHandler) ReceiveSyncPayload(c *gin.Context) {
	var job domain.SyncQueue

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado", "detalle": err.Error()})
		return
	}

	if err := h.service.QueueJob(c.Request.Context(), &job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Payload recibido y encolado para sincronización",
		"job_id":  job.ID,
	})
}
