package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type ProductAttributeValueHandler struct {
	service ports.ProductAttributeValueService
}

func NewProductAttributeValueHandler(service ports.ProductAttributeValueService) *ProductAttributeValueHandler {
	return &ProductAttributeValueHandler{service: service}
}

func (h *ProductAttributeValueHandler) CreateValue(c *gin.Context) {
	var pav domain.ProductAttributeValue
	if err := c.ShouldBindJSON(&pav); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	if err := h.service.CreateValue(c.Request.Context(), &pav); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pav)
}

func (h *ProductAttributeValueHandler) GetValueByID(c *gin.Context) {
	id := c.Param("id")
	pav, err := h.service.GetValueByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pav)
}
