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

// Crear un nuevo usuario
func (r *UsersRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// Editar un usuario por su id
func (r *UsersRepository) UpdateUser(id uint, user *models.User) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

// Eliminar un usuario
func (r *UsersRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// Retornar un usuario por su correo y dni
func (r *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("correo = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Cambiar contraseña de un usuario
func (r *UsersRepository) ChangePassword(id uint, newPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("dni", newPassword).Error
}
