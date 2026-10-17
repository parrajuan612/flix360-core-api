package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"flix360-core-api/internal/adapters/handlers/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRequireRole(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		userRole       interface{}
		hasRole        bool
		allowedRoles   []string
		expectedStatus int
	}{
		{
			name:           "Sin rol en el contexto retorna 403",
			userRole:       nil,
			hasRole:        false,
			allowedRoles:   []string{"admin"},
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "Rol insuficiente retorna 403",
			userRole:       "operator",
			hasRole:        true,
			allowedRoles:   []string{"admin"},
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "Rol admin permitido retorna 200",
			userRole:       "admin",
			hasRole:        true,
			allowedRoles:   []string{"admin"},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Case-insensitive rol Admin permitido",
			userRole:       "ADMIN",
			hasRole:        true,
			allowedRoles:   []string{"admin"},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Uno de múltiples roles permitidos",
			userRole:       "manager",
			hasRole:        true,
			allowedRoles:   []string{"admin", "manager"},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.Use(func(c *gin.Context) {
				if tt.hasRole {
					c.Set("role", tt.userRole)
				}
				c.Next()
			})
			router.GET("/test", middlewares.RequireRole(tt.allowedRoles...), func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("esperado status %d, obtenido %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
