package http

import (
	"net/http"
	"strconv"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// ProductHandler maneja las peticiones HTTP para los productos
type ProductHandler struct {
	service ports.ProductService
}

// NewProductHandler es el constructor que inyecta el servicio de negocio
func NewProductHandler(service ports.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// CreateProduct maneja el POST /api/v1/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product

	// 1. Gin mapea mágicamente el JSON entrante al Struct de Go
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos en el JSON",
			"detalle": err.Error(),
		})
		return
	}

	// 2. Pasamos el contexto y la información a nuestra capa de negocio (Servicio)
	// Aquí es donde brillan los principios SOLID: el handler no sabe NADA de validaciones o bases de datos.
	if err := h.service.CreateProduct(c.Request.Context(), &product); err != nil {
		// Si la validación de negocio falla (ej. falta el nombre), devolvemos un 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 3. Si todo salió bien, respondemos con un 201 Created y el objeto completo (incluyendo el ID autogenerado)
	c.JSON(http.StatusCreated, product)
}

// GetProduct maneja el GET /api/v1/products/:id
func (h *ProductHandler) GetProduct(c *gin.Context) {
	// 1. Extraemos el ID directamente de los parámetros de la URL
	id := c.Param("id")

	// 2. Solicitamos el producto al servicio
	product, err := h.service.GetProduct(c.Request.Context(), id)
	if err != nil {
		// Si hay error (ej. el repositorio dice que no existe), mandamos un 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Producto no encontrado",
			"detalle": err.Error(),
		})
		return
	}

	// 3. Respondemos con un 200 OK y el JSON del producto
	c.JSON(http.StatusOK, product)
}
func (h *ProductHandler) ListProducts(c *gin.Context) {
	// Extraemos la empresa del Token de seguridad
	companyID := c.GetString("company_id")

	// Leer parámetros de la URL (Query Params)
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")
	search := c.Query("search")

	// Convertir strings a enteros
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	// Llamar al servicio
	response, err := h.service.ListProducts(c.Request.Context(), companyID, limit, offset, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar productos"})
		return
	}

	c.JSON(http.StatusOK, response)
}
