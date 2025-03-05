package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
)

type UsersServices struct {
	UsersRepo *repository.UsersRepository
}

func NewUsersServices(usersrepo *repository.UsersRepository) *UsersServices {
	return &UsersServices{UsersRepo: usersrepo}
}

func (s *UsersServices) GetAllUsers() ([]models.User, error) {
	return s.UsersRepo.GetAll()
}
