package main

import (
	"barberia/internal/database/mysql"
	"barberia/internal/middleware"
	"barberia/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Crear instancia de Echo
	e := echo.New()

	// Aplicar middleware de logs a todas las rutas (middlewar global)
	e.Use(middleware.LoggerMiddleware)

	database := new(mysql.Config)
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Configurar todas las rutas
	routes.SetupRoutes(e, db)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}
