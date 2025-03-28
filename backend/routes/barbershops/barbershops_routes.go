package barbershops_routes

import (
	"barberia/internal/handlers"
	"barberia/internal/repository"
	"barberia/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRoutes(e *echo.Echo, db *gorm.DB, private *echo.Group) {
	barbershopRepo := repository.NewBarbershopsRepository(db)               // Crear un nuevo repositorio para las barberías
	barbershopServices := services.NewBarberShopsServices(barbershopRepo)   // Crear un nuevo servicio para las barberías
	barbershopHandler := handlers.NewBarberShopsHandler(barbershopServices) // Crear un manejador de solicitud para las barberías

	// Obtener (GET)
	private.GET("/barbershops", barbershopHandler.GetAllBarberShops)    // Obtener todas las barberías
	private.GET("/barbershop/:id", barbershopHandler.GetByIDBarbershop) // Obtener una barbería por su ID

	// Crear (POST)
	private.POST("/create/barbershop", barbershopHandler.CreateBarbershop) // Crear una nueva barbería

	// Eliminar (DELETE)
	private.DELETE("/delete/barbershop/:id", barbershopHandler.DeleteBarbershop) // Eliminar una barbería por su ID

	// Editar (PUT)
	private.PUT("/update/barbershop/:id", barbershopHandler.BarbershopUpdate) // Editar una barbería por su ID
}
