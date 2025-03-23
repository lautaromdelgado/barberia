package routes

import (
	auth_middleware "barberia/internal/middleware/auth"
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
	public := e.Group("/public")   // Rutas públicas
	private := e.Group("/private") // Rutas privadas

	// Middleware de autenticación para las rutas privadas
	private.Use(auth_middleware.AuthMiddleware)

	// Rutas privadas
	users_routes.SetupRoutes(e, db, private)               // Rutas de los usuarios
	barbershops_routes.SetUpRoutes(e, db, private)         // Rutas de las barberías
	haircuts_routes.SetUpRoutes(e, db, private)            // Rutas de los cortes
	barbershops_employee_routes.SetUpRoute(e, db, private) // Rutas de los empleados de las barberías

	// Rutas públicas
	auth_routes.SetUpRoutes(e, db, public) // Rutas de autenticación
}
