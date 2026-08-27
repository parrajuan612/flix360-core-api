package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	service ports.DeviceService
}

func NewDeviceHandler(service ports.DeviceService) *DeviceHandler {
	return &DeviceHandler{service: service}
}

func (h *DeviceHandler) CreateDevice(c *gin.Context) {
	var dev domain.Device

	if err := c.ShouldBindJSON(&dev); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	if err := h.service.CreateDevice(c.Request.Context(), &dev); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dev)
}

func (h *DeviceHandler) GetDeviceByID(c *gin.Context) {
	id := c.Param("id")

	dev, err := h.service.GetDeviceByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dev)
}
