package http

import (
	"net/http"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type RfidTagHandler struct {
	service ports.RfidTagService
}

func NewRfidTagHandler(service ports.RfidTagService) *RfidTagHandler {
	return &RfidTagHandler{service: service}
}

func (h *RfidTagHandler) CreateTag(c *gin.Context) {
	var tag domain.RfidTag

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	// 🔥 LA MAGIA DE SEGURIDAD AQUÍ: Inyectamos el ID de la empresa desde el Token
	tag.CompanyID = c.GetString("company_id")

	if err := h.service.CreateTag(c.Request.Context(), &tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

func (h *RfidTagHandler) GetTagByEPC(c *gin.Context) {
	epc := c.Param("epc")

	tag, err := h.service.GetTagByEPC(c.Request.Context(), epc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}
