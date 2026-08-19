package api

import (
	"fmt"
	"log"
	"os"

	handlerHttp "flix360-core-api/internal/adapters/handlers/http"
	"flix360-core-api/internal/adapters/repositories/postgres"
	"flix360-core-api/internal/core/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Driver de Postgres
)

func InitServer(r *gin.Engine) {
	// 1. Obtener cadena de conexión (DSN) del .env
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// DSN por defecto para entorno local (¡Ajusta tu usuario y contraseña aquí!)
		dsn = "postgres://postgres:password@localhost:5432/flix360?sslmode=disable"
	}

	// 2. Conectar a PostgreSQL usando sqlx
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error conectando a la base de datos: ", err)
	}
	log.Println("✅ Conexión exitosa a la base de datos PostgreSQL")

	// 3. Inyección de dependencias (Wiring)
	productRepo := postgres.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlerHttp.NewProductHandler(productService)

	// 4. Inicializar Rutas
	InitRouter(r, productHandler)

	// 5. Levantar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090" // Mantengo el puerto 8090 que mostraste en tu ejemplo
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	r.Run(fmt.Sprintf(":%s", port))
}
