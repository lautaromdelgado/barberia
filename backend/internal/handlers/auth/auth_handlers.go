package handlers

import (
	"barberia/internal/models"
	"barberia/internal/services"
	auth_services "barberia/internal/services/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandlers es una estructura que representa los manejadores de autenticación
type AuthHandlers struct {
	JWTService    *auth_services.JWTService
	UsersServices *services.UsersServices
}

// NewAuthHandlers es un constructor de AuthHandlers
func NewAuthHandlers(jwtService *auth_services.JWTService, usersServices *services.UsersServices) *AuthHandlers {
	return &AuthHandlers{
		JWTService:    jwtService,    // Servicio de JWT
		UsersServices: usersServices, // Servicio de usuarios
	}
}

// Register es un manejador que se encarga de registrar un usuario
func (a *AuthHandlers) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "invalid date: " + err.Error()})
	}
	if err := a.UsersServices.CreateNewUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "error registering a user",
		})
	}
	token, err := a.JWTService.GenerateToken(user.ID, user.Correo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "error generating token",
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"token":   token,
		"message": "user created",
	})
}
