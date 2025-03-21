package routes

import (
	auth_routes "barberia/routes/auth"
	barbershops_routes "barberia/routes/barbershops"
	barbershops_employee_routes "barberia/routes/barbershops_employee"
	haircuts_routes "barberia/routes/haircuts"
	users_routes "barberia/routes/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes inicializa todas las rutas de la aplicación
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	users_routes.SetupRoutes(e, db)               // Rutas de los usuarios
	barbershops_routes.SetUpRoutes(e, db)         // Rutas de las barberías
	haircuts_routes.SetUpRoutes(e, db)            // Rutas de los cortes
	barbershops_employee_routes.SetUpRoute(e, db) // Rutas de los empleados de las barberías
	auth_routes.SetUpRoutes(e, db)                // Rutas de autenticación
}
