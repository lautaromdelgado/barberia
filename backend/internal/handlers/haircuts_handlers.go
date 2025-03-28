package handlers

import (
	"barberia/internal/models"
	"barberia/internal/services"
	"net/http"

	"strconv"

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

// GetHaircutByID retorna un corte por su ID
func (h *HaircutsHandlers) GetHaircutByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	haircuts, err := h.HaircutsServices.GetHaircutByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, haircuts)
}

// Eliminar un corte por su ID
func (h *HaircutsHandlers) DeleteHaircut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	if err := h.HaircutsServices.DeleteHaircut(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "haircut deleted",
	})
}

// Actualizar un corte
func (h *HaircutsHandlers) UpdateHaircut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	var haircut models.Haircut
	if err := c.Bind(&haircut); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	if err := h.HaircutsServices.UpdateHaircut(&haircut, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "haircut updated successfully",
	})
}

// Crear un corte en la base de datos
func (h *HaircutsHandlers) CreateHaircut(c echo.Context) error {
	var haircut models.Haircut
	if err := c.Bind(&haircut); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	if err := h.HaircutsServices.CreateHaircut(&haircut); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"message": "haircut created successfully",
	})
}
