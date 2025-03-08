package handlers

import (
	"barberia/internal/models"
	"barberia/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Estructura de la clase UsersHandler
type UsersHandler struct {
	UsersServices *services.UsersServices // Servicios de usuarios
}

// Constructor de la clase UsersHandler
func NewUsersHandler(userservice *services.UsersServices) *UsersHandler {
	return &UsersHandler{UsersServices: userservice}
}

// Retornar todos los usuarios de la base de datos
func (h *UsersHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UsersServices.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

// Retornar un usuario por su id
func (h *UsersHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "Invalid id",
		})
	}
	user, err := h.UsersServices.GetUserById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}

// Crear un nuevo usuario
func (h *UsersHandler) CreateNewUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "Invalid data",
		})
	}
	if err := h.UsersServices.CreateNewUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"message": "User created successfully",
	})
}
