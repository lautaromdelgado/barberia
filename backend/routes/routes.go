package routes

import (
	barbershops_routes "barberia/routes/barbershops"
	users_routes "barberia/routes/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes inicializa todas las rutas de la aplicación
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	users_routes.SetupRoutes(e, db)       // Rutas de los usuarios
	barbershops_routes.SetUpRoutes(e, db) // Rutas de las barberías
}
