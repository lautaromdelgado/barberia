package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
)

// BarberShopsService es un servicio para los repositorios de BarberShops
type BarberShopsServices struct {
	BarberRepo *repository.Barbershops
}

// NewBarberShopsService crea un nuevo servicio para los repositorios de BarberShops
func NewBarberShopsServices(barberRepo *repository.Barbershops) *BarberShopsServices {
	return &BarberShopsServices{BarberRepo: barberRepo}
}

// GetAllBarberShops retorna todas las barberías de la base de datos
func (b *BarberShopsServices) GetAllBarberShops() ([]models.Barbershop, error) {
	return b.BarberRepo.GetAll()
}

// GetByIDBarberShop retorna una barbería por su ID
func (b *BarberShopsServices) GetByIDBarbershops(id uint) (*models.Barbershop, error) {
	return b.BarberRepo.GetByID(id)
}
