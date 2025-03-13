package handlers

import (
	"barberia/internal/models"
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

// Obtener todos los empleados de la base de datos
func (b *BarbershopEmployeeHandler) GetAllEmployees(c echo.Context) error {
	employees, totalEmployees, err := b.BarbershopEmployeeService.GetAllEmployees()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total_employees:": totalEmployees,
		"employees":        employees,
	})
}

// Borrar un empleado de la base de datos
func (b *BarbershopEmployeeHandler) DeleteEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	if err := b.BarbershopEmployeeService.DeleteEmployee(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "employee deleted successfully",
	})
}

// Crear un empleado en la base de datos
func (b *BarbershopEmployeeHandler) CreateEmployeee(c echo.Context) error {
	var employee models.BarbershopEmployee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid data",
		})
	}
	if err := b.BarbershopEmployeeService.CreateEmployee(&employee); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"message": "employee created successfully",
	})
}

// Actuaalizar un empleado
func (b *BarbershopEmployeeHandler) UpdateEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	var employee models.BarbershopEmployee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid data",
		})
	}
	if err := b.BarbershopEmployeeService.UpdateEmployee(&employee, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"message": "employee updated successfully",
	})
}
