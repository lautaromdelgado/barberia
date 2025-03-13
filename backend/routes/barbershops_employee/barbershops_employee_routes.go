package barbershops_employee_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRoute(e *echo.Echo, db *gorm.DB) {
	barbershopEmployeeRepo := repository.NewBarbershopEmployeeRepository(db)                      // Crear un nuevo repositorio de empleados de barberías
	barbershopEmployeeService := services.NewBarbershopEmployeeService(barbershopEmployeeRepo)    // Crear un nuevo servicio de empleados de barberías
	barbershopEmployeeHandler := handlers.NewBarbershopEmployeeHandler(barbershopEmployeeService) // Crear un nuevo manejador de empleados de barberías

	// Obtener (GET)
	e.GET("/employees/:id", barbershopEmployeeHandler.GetEmployeeByID) // Obtener datos de un empleado
	e.GET("/employees", barbershopEmployeeHandler.GetAllEmployees)     // Obtener todos los empleados registrados en la base de datos

	// Crear (POST)
	e.POST("/create/employees", barbershopEmployeeHandler.CreateEmployeee) // Crear un nuevo empleado

	// Actualizar (PUT)

	// Eliminar (DELETE)
	e.DELETE("/delete/employees/:id", barbershopEmployeeHandler.DeleteEmployee) // Borrar un empleado de la base de datos
}
