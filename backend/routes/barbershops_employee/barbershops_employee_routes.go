package barbershops_employee_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRoute(e *echo.Echo, db *gorm.DB) {
	barbershopEmployeeRepo := repository.NewBarbershopEmployeeRepository(db)
	barbershopEmployeeService := services.NewBarbershopEmployeeService(barbershopEmployeeRepo)
	barbershopEmployeeHandler := handlers.NewBarbershopEmployeeHandler(barbershopEmployeeService)

	// Obtener (GET)
	e.GET("/employees/:id", barbershopEmployeeHandler.GetEmployeeByID)
	e.GET("/employees", barbershopEmployeeHandler.GetAllEmployees)

	// Crear (POST)

	// Actualizar (PUT)

	// Eliminar (DELETE)
	e.DELETE("/delete/employees/:id", barbershopEmployeeHandler.DeleteEmployee)
}
