package users_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes inicializa las rutas de usuarios
func SetupRoutes(e *echo.Echo, db *gorm.DB, private *echo.Group) {
	usersRepo := repository.NewUsersRepository(db)
	usersService := services.NewUsersServices(usersRepo)
	usersHandler := handlers.NewUsersHandler(usersService)

	// Obtener (GET)
	private.GET("/users", usersHandler.GetAllUsers)    // Obtener todos los usuarios
	private.GET("/user/:id", usersHandler.GetUserById) // Obtener un usuario por su id

	// Crear (POST)
	private.POST("/create/user", usersHandler.CreateNewUser) // Crear un nuevo usuario

	// Editar (PUT)
	private.PUT("/update/user/:id", usersHandler.UpdateUser)              // Editar un usuario por su id
	private.PUT("/user/:id/change-password", usersHandler.ChangePassword) // Cambiar la contraseña de un usuario

	// Eliminar (DELETE)
	private.DELETE("/delete/user/:id", usersHandler.DeleteUser) // Eliminar un usuario
}
