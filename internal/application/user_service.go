package application

import (
	"templateApiRestGo/internal/domain"
)

// 🧠 Application Layer (Capa de aplicación):
// Aquí va la lógica de negocio y las interfaces (puertos).

// Puerto (interfaz) que el servicio necesita.
// Define qué operaciones pueden hacerse sobre los usuarios.
type UserRepository interface {
	Create(user *domain.User) error
}

// Servicio que usa la interfaz (puerto) para ejecutar lógica de negocio.
type UserService struct {
	repo UserRepository
}

// Constructor del servicio.
// Aquí se inyecta el repositorio (ya sea real o un mock para test).
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Caso de uso: Crear un usuario.
func (s *UserService) CreateUser(user *domain.User) error {
	// Aquí podría ir lógica extra (validaciones, eventos, etc.)
	return s.repo.Create(user)
}
