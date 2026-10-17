package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, companyID string, req *domain.CreateUserRequest) (*domain.User, error) {
	if strings.TrimSpace(companyID) == "" {
		return nil, errors.New("el company_id es obligatorio")
	}

	if len(req.Password) < 8 {
		return nil, errors.New("la contraseña debe tener al menos 8 caracteres")
	}

	// 1. Validar existencia previa de email
	existing, _ := s.repo.GetByEmail(ctx, strings.ToLower(strings.TrimSpace(req.Email)))
	if existing != nil {
		return nil, errors.New("el correo electrónico ya está registrado")
	}

	// 2. Hashear la contraseña con bcrypt (costo por defecto = 10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error al procesar la contraseña")
	}

	// 3. Asignar rol por defecto
	role := strings.ToLower(strings.TrimSpace(req.Role))
	if role == "" {
		role = "operator"
	}
	if role != "admin" && role != "manager" && role != "operator" {
		return nil, errors.New("rol inválido. Debe ser: admin, manager u operator")
	}

	now := time.Now().UTC()
	user := &domain.User{
		ID:           uuid.New().String(),
		CompanyID:    companyID, // Forzado desde el token del admin autenticado
		FullName:     strings.TrimSpace(req.FullName),
		Email:        strings.ToLower(strings.TrimSpace(req.Email)),
		PasswordHash: string(hashedPassword),
		Role:         role,
		Status:       "active",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Limpiar hash antes de retornar
	user.PasswordHash = ""
	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
func (s *userService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("el email no puede estar vacío")
	}
	return s.repo.GetByEmail(ctx, email)
}
