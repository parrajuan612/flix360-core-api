package api

import (
	"context" // <-- Nuevo import
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitServer(r *gin.Engine) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://flix_admin:flix_password@localhost:5433/flix360?sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error conectando a la base de datos: ", err)
	}
	log.Println("✅ Conexión exitosa a la base de datos PostgreSQL")

	// ---------------------------------------------------------
	// Magia Clean Code: Toda la inyección ocurre en una sola línea
	// ---------------------------------------------------------
	handlers := BuildDependencies(db)

	// ---------------------------------------------------------
	// ARRANQUE DEL MOTOR OFFLINE (Goroutine)
	// ---------------------------------------------------------
	go handlers.Worker.Start(context.Background())

	// Inicializar Rutas pasándole el contenedor completo
	InitRouter(r, handlers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8590"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Fallo al arrancar el servidor: %v", err)
	}
}
