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

type categoryService struct {
	repo ports.CategoryRepository
}

func NewCategoryService(repo ports.CategoryRepository) ports.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(ctx context.Context, cat *domain.Category) error {
	if strings.TrimSpace(cat.CompanyID) == "" {
		return errors.New("el company_id es obligatorio")
	}
	if strings.TrimSpace(cat.Name) == "" {
		return errors.New("el nombre de la categoría es obligatorio")
	}

	cat.ID = uuid.New().String()

	if cat.Status == "" {
		cat.Status = "active"
	}

	now := time.Now().UTC()
	cat.CreatedAt = now
	cat.UpdatedAt = now

	return s.repo.Create(ctx, cat)
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}
	return s.repo.GetByID(ctx, id)
}
func (s *categoryService) ListCategories(ctx context.Context, companyID string, limit, offset int) (*domain.PaginatedResponse, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	cats, total, err := s.repo.List(ctx, companyID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResponse{Total: total, Limit: limit, Offset: offset, Data: cats}, nil
}
