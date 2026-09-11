package services

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo ports.UserRepository
}

func NewAuthService(userRepo ports.UserRepository) ports.AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	// 1. Buscar al usuario
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("credenciales incorrectas")
	}

	// 2. Validar estado del usuario
	if user.Status != "active" {
		return nil, errors.New("el usuario está inactivo o bloqueado")
	}

	// 3. Comparar contraseñas
	// TODO: En un entorno de producción usaremos bcrypt.CompareHashAndPassword
	// Por ahora validamos el texto plano que insertamos en nuestra prueba
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("credenciales incorrectas")
	}

	// 4. Crear el Token (JWT)
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "super_secreto_para_desarrollo_local" // Fallback local
	}

	// Creamos el payload del token (los datos que van adentro)
	claims := jwt.MapClaims{
		"user_id":    user.ID,
		"company_id": user.CompanyID,
		"role":       user.Role,
		"exp":        time.Now().Add(time.Hour * 24 * 7).Unix(), // ¡Caduca en 7 días! (Modo offline)
		"iat":        time.Now().Unix(),                         // Fecha de creación
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, errors.New("error interno al generar el token")
	}

	// 5. Limpiar contraseña antes de enviarlo al cliente
	user.PasswordHash = ""

	return &domain.LoginResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}
