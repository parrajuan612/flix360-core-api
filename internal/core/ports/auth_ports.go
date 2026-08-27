package ports

import (
	"context"
	"flix360-core-api/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error)
}
