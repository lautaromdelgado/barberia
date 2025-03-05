package handlers

import (
	"barberia/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	UsersServices *services.UsersServices
}

func NewUsersHandler(userservice *services.UsersServices) *UsersHandler {
	return &UsersHandler{UsersServices: userservice}
}

func (h *UsersHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UsersServices.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}
