package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireRole valida que el rol en el token coincida con al menos uno de los roles permitidos
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: rol no identificado"})
			c.Abort()
			return
		}

		userRole := strings.ToLower(fmt.Sprintf("%v", roleVal))
		for _, allowed := range allowedRoles {
			if strings.EqualFold(userRole, allowed) {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: privilegios insuficientes"})
		c.Abort()
	}
}
