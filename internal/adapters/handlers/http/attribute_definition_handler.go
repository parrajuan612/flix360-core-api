package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type AttributeDefinitionHandler struct {
	service ports.AttributeDefinitionService
}

func NewAttributeDefinitionHandler(service ports.AttributeDefinitionService) *AttributeDefinitionHandler {
	return &AttributeDefinitionHandler{service: service}
}

func (h *AttributeDefinitionHandler) CreateAttribute(c *gin.Context) {
	var attr domain.AttributeDefinition
	if err := c.ShouldBindJSON(&attr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	if err := h.service.CreateAttribute(c.Request.Context(), &attr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attr)
}

func (h *AttributeDefinitionHandler) GetAttributeByID(c *gin.Context) {
	id := c.Param("id")
	attr, err := h.service.GetAttributeByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attr)
}
