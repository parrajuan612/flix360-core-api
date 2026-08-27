package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type InventoryMovementHandler struct {
	service ports.InventoryMovementService
}

func NewInventoryMovementHandler(service ports.InventoryMovementService) *InventoryMovementHandler {
	return &InventoryMovementHandler{service: service}
}

func (h *InventoryMovementHandler) CreateMovement(c *gin.Context) {
	var mov domain.InventoryMovement

	if err := c.ShouldBindJSON(&mov); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	if err := h.service.CreateMovement(c.Request.Context(), &mov); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mov)
}

func (h *InventoryMovementHandler) GetMovementByID(c *gin.Context) {
	id := c.Param("id")

	mov, err := h.service.GetMovementByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mov)
}
func (h *InventoryMovementHandler) CreateBulk(c *gin.Context) {
	// ¡Extraemos los datos seguros que dejó el Guardia (Middleware)!
	companyID := c.GetString("company_id")
	userID := c.GetString("user_id")

	var movs []*domain.InventoryMovement

	// Go espera recibir un JSON en formato de arreglo: [ {...}, {...} ]
	if err := c.ShouldBindJSON(&movs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de arreglo inválido", "detalle": err.Error()})
		return
	}

	if err := h.service.CreateBulk(c.Request.Context(), companyID, userID, movs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Movimientos registrados exitosamente",
		"total":   len(movs),
	})
}
