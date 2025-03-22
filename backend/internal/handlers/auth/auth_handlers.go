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
	token, err := a.JWTService.GenerateToken(user.ID, user.Nombre, user.Apellido, user.Correo, user.Rol)
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

// Login es un manejador que se encarga de loguear un usuario
func (a *AuthHandlers) Login(c echo.Context) error {
	type input struct {
		Correo string `json:"correo"`
		DNI    string `json:"dni"`
	}
	var inputData input
	if err := c.Bind(&inputData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "invalid data: " + err.Error(),
		})
	}
	token, err := a.UsersServices.GetUserByEmailAndDNI(inputData.Correo, inputData.DNI)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"status":  "error",
			"message": "error logging in: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"token":   token,
		"message": "logged in",
	})
}
