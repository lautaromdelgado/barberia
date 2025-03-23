package middleware

import (
	auth "barberia/internal/services/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtener el token del header de autorización
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"status":  "unauthorized",
				"message": "Authorization header is required",
			})
		}
		// Verificar que el token sea del formato correcto
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"status":  "unauthorized",
				"message": "invalid token format",
			})
		}
		// Verificar que el token sea válido
		claims, err := auth.NewJWTService().ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"status":  "unauthorized",
				"message": "invalid or expired token",
			})
		}
		// Guardar los claims en el contexto
		c.Set("user", claims)
		return next(c)
	}
}
