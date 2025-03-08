package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
	"errors"
)

// Estructura de la clase UsersServices
type UsersServices struct {
	UsersRepo *repository.UsersRepository // Repositorio de usuarios
}

// Constructor de la clase UsersServices
func NewUsersServices(usersrepo *repository.UsersRepository) *UsersServices {
	return &UsersServices{UsersRepo: usersrepo}
}

// Retornar todos los usuarios de la base de datos
func (s *UsersServices) GetAllUsers() ([]models.User, error) {
	return s.UsersRepo.GetAll()
}

// Retornar un usuario por su id
func (s *UsersServices) GetUserById(id uint) (*models.User, error) {
	user, err := s.UsersRepo.GetById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Crear un nuevo usuario
func (s *UsersServices) CreateNewUser(user *models.User) error {
	if user.Nombre == "" || user.Apellido == "" || user.DNI == "" {
		return errors.New("missing required fields")
	}
	return s.UsersRepo.CreateUser(user)
}

// Editar un usuario por su id
func (s *UsersServices) UpdateUser(id uint, user *models.User) error {
	if user.Nombre != "" && user.Apellido != "" && user.DNI != "" {
		return errors.New("at least one field is required")
	}
	return s.UsersRepo.UpdateUser(id, user)
}

// Eliminar un usuario
func (s *UsersServices) DeleteUser(id uint) error {
	return s.UsersRepo.DeleteUser(id)
}
