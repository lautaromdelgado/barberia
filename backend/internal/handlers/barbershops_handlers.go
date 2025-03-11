package handlers

import (
	"barberia/internal/models"
	"barberia/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// BarberShopsHandler es un manejador para los servicios de BarberShops
type BarberShopsHandler struct {
	BarberShopServices *services.BarberShopsServices
}

// NewBarberShopsHandler crea un nuevo manejador para los servicios de BarberShops
func NewBarberShopsHandler(barbershopsServices *services.BarberShopsServices) *BarberShopsHandler {
	return &BarberShopsHandler{BarberShopServices: barbershopsServices}
}

// GetAllBarberShops retorna todas las barberías
func (b *BarberShopsHandler) GetAllBarberShops(c echo.Context) error {
	barbershops, err := b.BarberShopServices.GetAllBarberShops()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, barbershops)
}

// GetByIDBarberShop retorna una barbería por su ID
func (b *BarberShopsHandler) GetByIDBarbershop(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid id",
		})
	}
	barbershop, err := b.BarberShopServices.GetByIDBarbershops(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "barbershop not found",
		})
	}
	return c.JSON(http.StatusOK, barbershop)
}

// CreateBarbershop crea una nueva barbería en la base de datos
func (b *BarberShopsHandler) CreateBarbershop(c echo.Context) error {
	var barbershop models.Barbershop
	if err := c.Bind(&barbershop); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "invalid data",
		})
	}
	if err := b.BarberShopServices.CreateBarbershop(&barbershop); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "barbershop created successfully",
	})
}
