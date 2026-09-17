package http

import (
	"net/http"
	"strconv"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type InventoryAssetHandler struct {
	service ports.InventoryAssetService
}

func NewInventoryAssetHandler(service ports.InventoryAssetService) *InventoryAssetHandler {
	return &InventoryAssetHandler{service: service}
}

func (h *InventoryAssetHandler) CreateAsset(c *gin.Context) {
	var asset domain.InventoryAsset

	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}
	asset.CompanyID = c.GetString("company_id")
	userID := c.GetString("user_id")
	asset.CreatedBy = &userID

	if err := h.service.CreateAsset(c.Request.Context(), &asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, asset)
}

func (h *InventoryAssetHandler) GetAssetByID(c *gin.Context) {
	id := c.Param("id")

	asset, err := h.service.GetAssetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
}
func (h *InventoryAssetHandler) ListAssets(c *gin.Context) {
	companyID := c.GetString("company_id")

	locationID := c.Query("location_id") // El frontend puede mandarlo o no
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	response, err := h.service.ListAssets(c.Request.Context(), companyID, locationID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar activos"})
		return
	}

	c.JSON(http.StatusOK, response)
}
