package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware verifica que la petición traiga un token válido
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtener el header de autorización
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización requerido"})
			c.Abort() // Detiene la petición aquí, no llega al controlador
			return
		}

		// 2. El formato debe ser "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido. Use 'Bearer <token>'"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "super_secreto_para_desarrollo_local"
		}

		// 3. Validar y desencriptar el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// 4. Extraer los datos y guardarlos en el contexto de la petición
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("company_id", claims["company_id"])
			c.Set("role", claims["role"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Estructura del token inválida"})
			c.Abort()
			return
		}

		// Si todo está bien, dejamos que la petición continúe hacia el controlador
		c.Next()
	}
}
