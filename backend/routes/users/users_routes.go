package users_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes inicializa las rutas de usuarios
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	usersRepo := repository.NewUsersRepository(db)
	usersService := services.NewUsersServices(usersRepo)
	usersHandler := handlers.NewUsersHandler(usersService)

	// Obtener (GET)
	e.GET("/users", usersHandler.GetAllUsers)    // Obtener todos los usuarios
	e.GET("/user/:id", usersHandler.GetUserById) // Obtener un usuario por su id

	// Crear (POST)
	e.POST("/create/user", usersHandler.CreateNewUser) // Crear un nuevo usuario

	// Editar (PUT)
	e.PUT("/update/user/:id", usersHandler.UpdateUser) // Editar un usuario por su id

	// Eliminar (DELETE)
	e.DELETE("/delete/user/:id", usersHandler.DeleteUser) // Eliminar un usuario
}
