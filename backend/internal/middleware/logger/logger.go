package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// LoggerMiddleware registra cada solicitud que llega al servidor
func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		log.Printf("[%s] %s %s - %v", c.Request().Method, c.Path(), c.RealIP(), duration)
		return err
	}
}
