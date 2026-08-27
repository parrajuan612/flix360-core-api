package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service ports.LocationService
}

func NewLocationHandler(service ports.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var loc domain.Location

	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	if err := h.service.CreateLocation(c.Request.Context(), &loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, loc)
}

func (h *LocationHandler) GetLocationByID(c *gin.Context) {
	id := c.Param("id")

	loc, err := h.service.GetLocationByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loc)
}
