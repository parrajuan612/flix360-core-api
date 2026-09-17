package api

import (
	"flix360-core-api/internal/adapters/handlers/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, h *AppHandlers) {
	v1 := r.Group("/api/v1")
	{
		// ==========================================
		// RUTAS PÚBLICAS (No requieren Token)
		// ==========================================
		v1.POST("/login", h.Auth.Login)
		v1.POST("/users", h.User.CreateUser) // Solo por ahora para poder crear usuarios

		// ==========================================
		// RUTAS PRIVADAS (Protegidas por JWT)
		// ==========================================
		protected := v1.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware()) // <-- Aquí activamos el Guardia
		{
			// --- Productos ---
			protected.POST("/products", h.Product.CreateProduct)
			protected.GET("/products", h.Product.ListProducts)
			protected.GET("/products/:id", h.Product.GetProduct)

			// --- Etiquetas RFID ---
			protected.POST("/rfid-tags", h.RfidTag.CreateTag)
			protected.GET("/rfid-tags/:epc", h.RfidTag.GetTagByEPC)

			// --- Locaciones ---
			protected.POST("/locations", h.Location.CreateLocation)
			protected.GET("/locations/:id", h.Location.GetLocationByID)
			protected.GET("/locations", h.Location.ListLocations)

			// --- Dispositivos ---
			protected.POST("/devices", h.Device.CreateDevice)
			protected.GET("/devices/:id", h.Device.GetDeviceByID)

			// --- Categorías ---
			protected.POST("/categories", h.Category.CreateCategory)
			protected.GET("/categories", h.Category.ListCategories)
			protected.GET("/categories/:id", h.Category.GetCategoryByID)

			// --- Definición de Atributos ---
			protected.POST("/attribute-definitions", h.AttributeDefinition.CreateAttribute)
			protected.GET("/attribute-definitions/:id", h.AttributeDefinition.GetAttributeByID)

			// --- Valores de Atributos de Producto ---
			protected.POST("/product-attribute-values", h.ProductAttributeValue.CreateValue)
			protected.GET("/product-attribute-values/:id", h.ProductAttributeValue.GetValueByID)

			// --- Activos de Inventario y Sync ---
			protected.POST("/inventory-assets", h.Asset.CreateAsset)
			protected.GET("/inventory-assets", h.Asset.ListAssets)
			protected.GET("/inventory-assets/:id", h.Asset.GetAssetByID)
			protected.POST("/sync-queue", h.SyncQueue.ReceiveSyncPayload)

			// --- Movimientos (Auditoría) ---
			// 1. Primero la ruta masiva específica (/bulk)
			protected.POST("/inventory-movements/bulk", h.Movement.CreateBulk)

			// 2. Luego las rutas individuales (raíz y comodín /:id)
			protected.POST("/inventory-movements", h.Movement.CreateMovement)
			protected.GET("/inventory-movements/:id", h.Movement.GetMovementByID)
		}
	}
}
