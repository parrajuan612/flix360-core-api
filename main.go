package main

import (
	"log"
	"os"

	"flix360-core-api/cmd/api" // Importamos tu paquete api

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("No se pudo cargar el archivo .env, usando variables del sistema:", err)
		}
	}
}

func main() {
	// Instancia por defecto de Gin con Logger y Recovery middleware
	r := gin.Default()

	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// Iniciar servidor delegando la lógica
	api.InitServer(r)
}
