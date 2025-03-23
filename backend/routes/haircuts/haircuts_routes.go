package haircuts_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetUpRoutes configura las rutas de cortes
func SetUpRoutes(e *echo.Echo, db *gorm.DB, private *echo.Group) {
	haircutsRepo := repository.NewHaircutsRepository(db)               // Crear repositorio de cortes
	haircutsServices := services.NewHaircutsServices(haircutsRepo)     // Crear servicio de cortes
	haircutsHandlers := handlers.NewHaircutsHandlers(haircutsServices) // Crear manejador de cortes

	// Obtener (GET)
	private.GET("/haircuts", haircutsHandlers.GetAllHaircuts)     // Obtener todos los cortes
	private.GET("/haircuts/:id", haircutsHandlers.GetHaircutByID) // Obtener un corte por su ID

	// Eliminar (DELETE)
	private.DELETE("/delete/haircuts/:id", haircutsHandlers.DeleteHaircut) // Eliminar un corte por su ID

	// Actualizar (PUT)
	private.PUT("/update/haircuts/:id", haircutsHandlers.UpdateHaircut) // Actualizar un corte

	// Crear (POST)
	private.POST("/create/haircuts", haircutsHandlers.CreateHaircut) // Crear un corte
}
