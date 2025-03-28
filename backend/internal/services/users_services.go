package services

import (
	"barberia/internal/dto"
	"barberia/internal/models"
	"barberia/internal/repository"
	auth_services "barberia/internal/services/auth"
	"barberia/internal/utils"
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
	hashedPass, err := utils.HashedPassword(user.DNI) // Pasamos el DNI como contraseña para hashear
	if err != nil {
		return errors.New("error hashing password")
	}
	user.DNI = hashedPass // Seteamos el DNI(hasheado como contraseña) al modelo
	return s.UsersRepo.CreateUser(user)
}

// Editar un usuario por su id
func (s *UsersServices) UpdateUser(id uint, user *models.User) error {
	if user.Nombre != "" && user.Apellido != "" && user.DNI != "" {
		return errors.New("at least one field is required")
	}
	if !user.Verified {
		return errors.New("invalid verified field")
	}
	return s.UsersRepo.UpdateUser(id, user)
}

// Eliminar un usuario
func (s *UsersServices) DeleteUser(id uint) error {
	return s.UsersRepo.DeleteUser(id)
}

// Retornar un usuario por su correo y dni
func (s *UsersServices) GetUserByEmailAndDNI(email, dni string) (string, error) {
	user, err := s.UsersRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	if err := utils.CheckPasswordHash(user.DNI, dni); err != nil {
		return "", errors.New("invalid password")
	}
	token, err := auth_services.NewJWTService().GenerateToken(user.ID, user.Nombre, user.Apellido, user.Correo, user.Rol)
	if err != nil {
		return "", errors.New("error generating token")
	}
	return token, nil
}

// Cambiar la contraseña de un usuario
func (s *UsersServices) ChangePassword(id uint, newPassword *dto.ChangePassword) error {
	// Corroboramos que la contraseña en ningún caso este vacía|
	if newPassword.CurrentPassword == "" || newPassword.NewPassword == "" || newPassword.ConfirmPassword == "" {
		return errors.New("missing required fields")
	}
	// Si la contraseña nueva y la confirmación no son iguales, retornamos un error
	if newPassword.NewPassword != newPassword.ConfirmPassword {
		return errors.New("passwords do not match")
	}
	// Obtener el usuario por id para comparar las contraseñas
	user, err := s.UsersRepo.GetById(id)
	if err != nil {
		return errors.New("user not found")
	}
	// Verificamos que la nueva contraseña y la confirmación sean iguales
	err = utils.CheckPasswordHash(user.DNI, newPassword.CurrentPassword)
	if err != nil {
		return errors.New("invalid current password")
	}
	// Verificamos que la nueva contraseña no sea igual a la actual
	if newPassword.CurrentPassword == newPassword.NewPassword {
		return errors.New("new password cannot be the same as current password")
	}
	// Hasheo la contraseña nueva
	hashedPassword, err := utils.HashedPassword(newPassword.NewPassword)
	if err != nil {
		return errors.New("error hashing password")
	}
	if err := s.UsersRepo.ChangePassword(id, hashedPassword); err != nil {
		return errors.New("error changing password")
	}
	return nil
}
