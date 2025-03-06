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
