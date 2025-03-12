package handlers

import (
	"barberia/internal/services"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// BarbershopEmployeeHandler es una estruct
type BarbershopEmployeeHandler struct {
	BarbershopEmployeeService *services.BarbershopEmployeeService
}

// NewBarbershopEmployeeHandler crea un nuevo manejador de empleados de barberías
func NewBarbershopEmployeeHandler(barbershopEmployeeService *services.BarbershopEmployeeService) *BarbershopEmployeeHandler {
	return &BarbershopEmployeeHandler{BarbershopEmployeeService: barbershopEmployeeService}
}

// Obtener datos de un empleado
func (b *BarbershopEmployeeHandler) GetEmployeeByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	employee, err := b.BarbershopEmployeeService.GetEmployeeByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, employee)
}
