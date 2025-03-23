package barbershops_employee_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRoute(e *echo.Echo, db *gorm.DB, private *echo.Group) {
	barbershopEmployeeRepo := repository.NewBarbershopEmployeeRepository(db)                      // Crear un nuevo repositorio de empleados de barberías
	barbershopEmployeeService := services.NewBarbershopEmployeeService(barbershopEmployeeRepo)    // Crear un nuevo servicio de empleados de barberías
	barbershopEmployeeHandler := handlers.NewBarbershopEmployeeHandler(barbershopEmployeeService) // Crear un nuevo manejador de empleados de barberías

	// Obtener (GET)
	private.GET("/employees/:id", barbershopEmployeeHandler.GetEmployeeByID) // Obtener datos de un empleado
	private.GET("/employees", barbershopEmployeeHandler.GetAllEmployees)     // Obtener todos los empleados registrados en la base de datos

	// Crear (POST)
	private.POST("/create/employees", barbershopEmployeeHandler.CreateEmployeee) // Crear un nuevo empleado

	// Actualizar (PUT)
	private.PUT("/update/employees/:id", barbershopEmployeeHandler.UpdateEmployee) // Actualizar un empleado de la base de datos

	// Eliminar (DELETE)
	private.DELETE("/delete/employees/:id", barbershopEmployeeHandler.DeleteEmployee) // Borrar un empleado de la base de datos
}
