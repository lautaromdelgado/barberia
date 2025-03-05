package routes

import (
	users_routes "barberia/routes/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	users_routes.SetupRoutes(e, db)
}
