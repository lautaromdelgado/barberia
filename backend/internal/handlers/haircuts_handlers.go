package handlers

import (
	"barberia/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HaircutsHandlers es el manejador de cortes
type HaircutsHandlers struct {
	HaircutsServices *services.HaircutsServices
}

// NewHaircutsHandlers es el constructor de HaircutsHandlers
func NewHaircutsHandlers(haircutsServices *services.HaircutsServices) *HaircutsHandlers {
	return &HaircutsHandlers{
		HaircutsServices: haircutsServices,
	}
}

// GetAllHaircuts retorna todos los cortes
func (h *HaircutsHandlers) GetAllHaircuts(c echo.Context) error {
	haircuts, err := h.HaircutsServices.GetAllHaircuts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, haircuts)
}
