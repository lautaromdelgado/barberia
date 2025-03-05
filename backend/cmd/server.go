package main

import (
	"barberia/internal/database/mysql"
	"barberia/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Crear instancia de Echo
	e := echo.New()

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
