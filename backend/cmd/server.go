package main

import (
	"fmt"
	"log"

	VariablesEntorno "barberia/config/variables_entorno"

	"github.com/labstack/echo/v4"
)

func main() {
	// Crear instancia de Config
	configEnv := new(VariablesEntorno.Config)
	// Cargar configuración una sola vez
	config, err := configEnv.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Imprimir DSN generado
	fmt.Println("DSN generado:", config.GetDNS())

	// Crear instancia de Echo
	e := echo.New()

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}
