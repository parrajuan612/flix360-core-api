package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	handlerHttp "flix360-core-api/internal/adapters/handlers/http"
	"flix360-core-api/internal/adapters/handlers/middlewares"
	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"
	"flix360-core-api/internal/core/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// mockUserRepo implementa ports.UserRepository en memoria
type mockUserRepo struct {
	users map[string]*domain.User
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{
		users: make(map[string]*domain.User),
	}
}

func (m *mockUserRepo) Create(ctx context.Context, u *domain.User) error {
	copied := *u
	m.users[u.Email] = &copied
	return nil
}

func (m *mockUserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("usuario no encontrado")
}

func (m *mockUserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	if u, ok := m.users[email]; ok {
		return u, nil
	}
	return nil, errors.New("credenciales incorrectas")
}

var _ ports.UserRepository = (*mockUserRepo)(nil)

func generateTestToken(userID, companyID, role string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "super_secreto_para_desarrollo_local"
	}

	claims := jwt.MapClaims{
		"user_id":    userID,
		"company_id": companyID,
		"role":       role,
		"exp":        time.Now().Add(time.Hour).Unix(),
		"iat":        time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func setupTestRouter(repo ports.UserRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	userService := services.NewUserService(repo)
	authService := services.NewAuthService(repo)

	userHandler := handlerHttp.NewUserHandler(userService)
	authHandler := handlerHttp.NewAuthHandler(authService)

	v1 := r.Group("/api/v1")
	{
		// Rutas públicas
		v1.POST("/login", authHandler.Login)

		// Rutas privadas
		protected := v1.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		{
			protected.POST("/users", middlewares.RequireRole("admin"), userHandler.CreateUser)
			protected.GET("/users/:id", userHandler.GetUserByID)
		}
	}

	return r
}

// 1. Criterio de Aceptación: Test de Bloqueo Público (sin header -> 401)
func TestAcceptance_PublicAccessBlocked(t *testing.T) {
	repo := newMockUserRepo()
	router := setupTestRouter(repo)

	payload := map[string]interface{}{
		"full_name": "Invasor Anonimo",
		"email":     "hacker@example.com",
		"password":  "password123",
		"role":      "admin",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("esperado status %d 401 Unauthorized, obtenido %d", http.StatusUnauthorized, w.Code)
	}
}

// 2. Criterio de Aceptación: Test de Rol Insuficiente (token operator -> 403)
func TestAcceptance_InsufficientRoleForbidden(t *testing.T) {
	repo := newMockUserRepo()
	router := setupTestRouter(repo)

	operatorToken, err := generateTestToken("user-op-1", "company-100", "operator")
	if err != nil {
		t.Fatalf("error al generar token: %v", err)
	}

	payload := map[string]interface{}{
		"full_name": "Nuevo Operador",
		"email":     "nuevo@example.com",
		"password":  "password123",
		"role":      "operator",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+operatorToken)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("esperado status %d 403 Forbidden, obtenido %d", http.StatusForbidden, w.Code)
	}
}

// 3. Criterio de Aceptación: Test de Creación Exitosa (token admin -> 201 y password_hash limpio)
func TestAcceptance_AdminCreateUserSuccess(t *testing.T) {
	repo := newMockUserRepo()
	router := setupTestRouter(repo)

	adminToken, err := generateTestToken("admin-1", "tenant-alpha", "admin")
	if err != nil {
		t.Fatalf("error al generar token: %v", err)
	}

	payload := map[string]interface{}{
		"full_name": "Carlos Gomez",
		"email":     "carlos@tenantalpha.com",
		"password":  "MiPasswordSeguro123",
		"role":      "operator",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("esperado status %d 201 Created, obtenido %d. Body: %s", http.StatusCreated, w.Code, w.Body.String())
	}

	var res domain.User
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("error parseando respuesta json: %v", err)
	}

	if res.CompanyID != "tenant-alpha" {
		t.Errorf("company_id debe coincidir con el token ('tenant-alpha'), obtenido '%s'", res.CompanyID)
	}
	if res.Email != "carlos@tenantalpha.com" {
		t.Errorf("email esperado 'carlos@tenantalpha.com', obtenido '%s'", res.Email)
	}
	if res.PasswordHash != "" {
		t.Errorf("password_hash debe estar vacío en el response, obtenido '%s'", res.PasswordHash)
	}
}

// 4. Criterio de Aceptación: Test de Login Funcional (creado por admin -> login con password plana exitoso)
func TestAcceptance_LoginWithNewlyCreatedUser(t *testing.T) {
	repo := newMockUserRepo()
	router := setupTestRouter(repo)

	// Paso A: Admin crea el usuario
	adminToken, err := generateTestToken("admin-1", "tenant-alpha", "admin")
	if err != nil {
		t.Fatalf("error generando token de admin: %v", err)
	}

	userEmail := "lucia@tenantalpha.com"
	userPlainPassword := "LuciaPassword2026!"

	createPayload := map[string]interface{}{
		"full_name": "Lucia Mendez",
		"email":     userEmail,
		"password":  userPlainPassword,
		"role":      "manager",
	}
	createBody, _ := json.Marshal(createPayload)

	createReq, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer "+adminToken)
	createW := httptest.NewRecorder()

	router.ServeHTTP(createW, createReq)

	if createW.Code != http.StatusCreated {
		t.Fatalf("falló la creación del usuario. Status: %d, Body: %s", createW.Code, createW.Body.String())
	}

	// Paso B: El usuario recién creado hace login usando su email y password en texto plano
	loginPayload := map[string]interface{}{
		"email":    userEmail,
		"password": userPlainPassword,
	}
	loginBody, _ := json.Marshal(loginPayload)

	loginReq, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()

	router.ServeHTTP(loginW, loginReq)

	if loginW.Code != http.StatusOK {
		t.Fatalf("esperado status %d 200 OK en login, obtenido %d. Body: %s", http.StatusOK, loginW.Code, loginW.Body.String())
	}

	var loginResp domain.LoginResponse
	if err := json.Unmarshal(loginW.Body.Bytes(), &loginResp); err != nil {
		t.Fatalf("error al deserializar respuesta de login: %v", err)
	}

	if loginResp.Token == "" {
		t.Fatal("el token devuelto en login no puede estar vacío")
	}
	if loginResp.User.Email != userEmail {
		t.Errorf("email de usuario en login esperado '%s', obtenido '%s'", userEmail, loginResp.User.Email)
	}
	if loginResp.User.PasswordHash != "" {
		t.Errorf("password_hash en LoginResponse debe estar vacío, obtenido '%s'", loginResp.User.PasswordHash)
	}
}
