package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"

	"github.com/google/uuid"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, u *domain.User) error {
	if strings.TrimSpace(u.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(u.Email) == "" || strings.TrimSpace(u.FullName) == "" {
		return errors.New("el nombre y el email son obligatorios")
	}

	u.ID = uuid.New().String()

	// Valores por defecto
	if u.Role == "" {
		u.Role = "operator" // admin, manager, operator
	}
	if u.Status == "" {
		u.Status = "active"
	}

	// Nota: En un sistema real, aquí encriptaríamos el password con bcrypt antes de guardarlo.
	if u.PasswordHash == "" {
		u.PasswordHash = "hashed_temporal_password"
	}

	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now

	return s.repo.Create(ctx, u)
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
