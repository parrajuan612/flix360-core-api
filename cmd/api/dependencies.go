package api

import (
	"github.com/jmoiron/sqlx"

	handlerHttp "flix360-core-api/internal/adapters/handlers/http"
	"flix360-core-api/internal/adapters/repositories/postgres"
	attrRepo "flix360-core-api/internal/adapters/repositories/postgres/attribute_definition"
	catRepo "flix360-core-api/internal/adapters/repositories/postgres/category"
	devRepo "flix360-core-api/internal/adapters/repositories/postgres/device"
	assetRepo "flix360-core-api/internal/adapters/repositories/postgres/inventory_asset"
	movRepo "flix360-core-api/internal/adapters/repositories/postgres/inventory_movement"
	locRepo "flix360-core-api/internal/adapters/repositories/postgres/location"
	pavRepo "flix360-core-api/internal/adapters/repositories/postgres/product_attribute_value"
	rfidRepo "flix360-core-api/internal/adapters/repositories/postgres/rfid_tag"
	syncRepo "flix360-core-api/internal/adapters/repositories/postgres/sync_queue"
	usrRepo "flix360-core-api/internal/adapters/repositories/postgres/user"
	"flix360-core-api/internal/core/services"
	"flix360-core-api/internal/workers" // <-- Nuevo import del Worker
)

type AppHandlers struct {
	Product               *handlerHttp.ProductHandler
	RfidTag               *handlerHttp.RfidTagHandler
	Asset                 *handlerHttp.InventoryAssetHandler
	Location              *handlerHttp.LocationHandler
	Movement              *handlerHttp.InventoryMovementHandler
	Device                *handlerHttp.DeviceHandler
	User                  *handlerHttp.UserHandler
	SyncQueue             *handlerHttp.SyncQueueHandler
	Category              *handlerHttp.CategoryHandler
	AttributeDefinition   *handlerHttp.AttributeDefinitionHandler
	ProductAttributeValue *handlerHttp.ProductAttributeValueHandler
	Auth                  *handlerHttp.AuthHandler
	Worker                *workers.SyncWorker // <-- Agregamos el Worker a nuestra caja de herramientas
}

func BuildDependencies(db *sqlx.DB) *AppHandlers {
	// Repositorios
	productRepo := postgres.NewProductRepository(db)
	rfidRepository := rfidRepo.NewRfidTagRepository(db)
	assetRepository := assetRepo.NewInventoryAssetRepository(db)
	locationRepository := locRepo.NewLocationRepository(db)
	movementRepository := movRepo.NewInventoryMovementRepository(db)
	deviceRepository := devRepo.NewDeviceRepository(db)
	userRepository := usrRepo.NewUserRepository(db)
	syncQueueRepository := syncRepo.NewSyncQueueRepository(db)
	categoryRepository := catRepo.NewCategoryRepository(db)
	attributeRepository := attrRepo.NewAttributeDefinitionRepository(db)
	pavRepository := pavRepo.NewProductAttributeValueRepository(db)

	// Servicios
	productService := services.NewProductService(productRepo)
	rfidService := services.NewRfidTagService(rfidRepository)
	assetService := services.NewInventoryAssetService(assetRepository)
	locationService := services.NewLocationService(locationRepository)
	movementService := services.NewInventoryMovementService(movementRepository)
	deviceService := services.NewDeviceService(deviceRepository)
	userService := services.NewUserService(userRepository)
	syncQueueService := services.NewSyncQueueService(syncQueueRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	attributeService := services.NewAttributeDefinitionService(attributeRepository)
	pavService := services.NewProductAttributeValueService(pavRepository)

	// Servicio de Autenticación
	authService := services.NewAuthService(userRepository)

	// Inyectamos las dependencias al Worker
	syncWorker := workers.NewSyncWorker(syncQueueRepository, movementService) // <-- Instancia Worker

	// Handlers
	return &AppHandlers{
		Product:               handlerHttp.NewProductHandler(productService),
		RfidTag:               handlerHttp.NewRfidTagHandler(rfidService),
		Asset:                 handlerHttp.NewInventoryAssetHandler(assetService),
		Location:              handlerHttp.NewLocationHandler(locationService),
		Movement:              handlerHttp.NewInventoryMovementHandler(movementService),
		Device:                handlerHttp.NewDeviceHandler(deviceService),
		User:                  handlerHttp.NewUserHandler(userService),
		SyncQueue:             handlerHttp.NewSyncQueueHandler(syncQueueService),
		Category:              handlerHttp.NewCategoryHandler(categoryService),
		AttributeDefinition:   handlerHttp.NewAttributeDefinitionHandler(attributeService),
		ProductAttributeValue: handlerHttp.NewProductAttributeValueHandler(pavService),
		Auth:                  handlerHttp.NewAuthHandler(authService),
		Worker:                syncWorker, // <-- Lo pasamos al servidor
	}
}
