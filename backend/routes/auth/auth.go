package auth_routes

import (
	handlers "barberia/internal/handlers/auth"
	"barberia/internal/repository"
	"barberia/internal/services"
	auth_services "barberia/internal/services/auth"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRoutes(e *echo.Echo, db *gorm.DB, public *echo.Group) {
	// Uso de la clase UsersRepository (Para crear el usuario)
	usersRepo := repository.NewUsersRepository(db)
	// Uso de la clase UsersServices (Para la lógica de negocio en el registro de usuario)
	usersService := services.NewUsersServices(usersRepo)
	// Uso de la clase JWTService (Para la generación de tokens)
	jwtServices := auth_services.NewJWTService()
	// Uso de la clase AuthHandlers (Para el manejo de la autenticación)
	authHandlers := handlers.NewAuthHandlers(jwtServices, usersService)

	// Rutas de autenticación
	public.POST("/register", authHandlers.Register) // Registrar un usuario
	public.POST("/login", authHandlers.Login)       // Loguear un usuario
}
