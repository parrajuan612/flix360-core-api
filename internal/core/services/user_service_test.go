package services_test

import (
	"context"
	"errors"
	"testing"

	"flix360-core-api/internal/core/domain"
	"flix360-core-api/internal/core/ports"
	"flix360-core-api/internal/core/services"

	"golang.org/x/crypto/bcrypt"
)

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

func TestCreateUser_Success(t *testing.T) {
	repo := newMockUserRepo()
	svc := services.NewUserService(repo)

	req := &domain.CreateUserRequest{
		FullName: "Juan Perez",
		Email:    "juan@example.com",
		Password: "password123",
		Role:     "admin",
	}

	user, err := svc.CreateUser(context.Background(), "company-123", req)
	if err != nil {
		t.Fatalf("se esperaba nil error, obtenido %v", err)
	}

	if user.CompanyID != "company-123" {
		t.Errorf("companyID esperado 'company-123', obtenido '%s'", user.CompanyID)
	}
	if user.PasswordHash != "" {
		t.Errorf("password_hash devuelto debe estar vacío, obtenido '%s'", user.PasswordHash)
	}
	if user.Role != "admin" {
		t.Errorf("rol esperado 'admin', obtenido '%s'", user.Role)
	}

	// Verificar que en repo se guardó con hash bcrypt válido
	saved := repo.users["juan@example.com"]
	if saved == nil {
		t.Fatal("el usuario debe estar guardado en el repositorio")
	}
	err = bcrypt.CompareHashAndPassword([]byte(saved.PasswordHash), []byte("password123"))
	if err != nil {
		t.Errorf("la contraseña guardada no coincide con el hash bcrypt: %v", err)
	}
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	repo := newMockUserRepo()
	repo.users["juan@example.com"] = &domain.User{
		ID:    "existing-id",
		Email: "juan@example.com",
	}

	svc := services.NewUserService(repo)
	req := &domain.CreateUserRequest{
		FullName: "Juan Perez",
		Email:    "JUAN@example.com",
		Password: "password123",
	}

	_, err := svc.CreateUser(context.Background(), "company-123", req)
	if err == nil {
		t.Fatal("se esperaba error de email duplicado, obtenido nil")
	}
	if err.Error() != "el correo electrónico ya está registrado" {
		t.Errorf("mensaje de error esperado 'el correo electrónico ya está registrado', obtenido '%s'", err.Error())
	}
}

func TestCreateUser_ValidationErrors(t *testing.T) {
	repo := newMockUserRepo()
	svc := services.NewUserService(repo)

	// Falta company_id
	_, err := svc.CreateUser(context.Background(), "  ", &domain.CreateUserRequest{
		FullName: "Test",
		Email:    "test@test.com",
		Password: "password123",
	})
	if err == nil || err.Error() != "el company_id es obligatorio" {
		t.Errorf("error inesperado para company_id vacío: %v", err)
	}

	// Password muy corta
	_, err = svc.CreateUser(context.Background(), "company-123", &domain.CreateUserRequest{
		FullName: "Test",
		Email:    "test@test.com",
		Password: "short",
	})
	if err == nil || err.Error() != "la contraseña debe tener al menos 8 caracteres" {
		t.Errorf("error inesperado para password corta: %v", err)
	}

	// Rol inválido
	_, err = svc.CreateUser(context.Background(), "company-123", &domain.CreateUserRequest{
		FullName: "Test",
		Email:    "test@test.com",
		Password: "password123",
		Role:     "superuser",
	})
	if err == nil || err.Error() != "rol inválido. Debe ser: admin, manager u operator" {
		t.Errorf("error inesperado para rol inválido: %v", err)
	}
}
