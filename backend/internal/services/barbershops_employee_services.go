package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
)

// BarbershopEmployeeService es una estructura que define un servicio de empleados de barberías
type BarbershopEmployeeService struct {
	BarbershopEmployeeRepo *repository.BarbershopEmployeeRepository
}

// NewBarbershopEmployeeService crea un nuevo servicio de empleados de barberías
func NewBarbershopEmployeeService(barbershopEmployeeRepo *repository.BarbershopEmployeeRepository) *BarbershopEmployeeService {
	return &BarbershopEmployeeService{BarbershopEmployeeRepo: barbershopEmployeeRepo}
}

// Obtener datos de un empleado
func (b *BarbershopEmployeeService) GetEmployeeByID(id uint) (*models.BarbershopEmployee, error) {
	return b.BarbershopEmployeeRepo.GetByID(id)
}
