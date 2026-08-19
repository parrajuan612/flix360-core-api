package api

import (
	handlerHttp "flix360-core-api/internal/adapters/handlers/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, productHandler *handlerHttp.ProductHandler) {
	v1 := r.Group("/api/v1")
	{
		// Rutas para el dominio de Productos
		v1.POST("/products", productHandler.CreateProduct)
		v1.GET("/products/:id", productHandler.GetProduct)
		v1.POST("/rfid/:tag", handlerHttp.RfidTagHandler)
	}
}
