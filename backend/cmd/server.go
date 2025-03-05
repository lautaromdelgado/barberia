package main

import (
	"barberia/internal/database/mysql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	database := new(mysql.Config)
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DB: %v\n", db)

	// Crear instancia de Echo
	e := echo.New()

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}
