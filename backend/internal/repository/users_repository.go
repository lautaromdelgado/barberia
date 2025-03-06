package repository

import (
	"barberia/internal/models"

	"gorm.io/gorm"
)

// Estructura de la clase UsersRepository
type UsersRepository struct {
	db *gorm.DB // Base de datos
}

// Constructor de la clase UsersRepository
func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Retonar todos los usuarios de la base de datos
func (r *UsersRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Retornar un usuario por su id
func (r *UsersRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
