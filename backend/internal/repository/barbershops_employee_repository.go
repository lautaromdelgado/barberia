package repository

import (
	"barberia/internal/models"

	"gorm.io/gorm"
)

// BarbershopEmployeeRepository es una estructura que define un repositorio de empleados de barberías
type BarbershopEmployeeRepository struct {
	db *gorm.DB
}

// NewBarbershopEmployeeRepository crea un nuevo repositorio de empleados de barberías
func NewBarbershopEmployeeRepository(db *gorm.DB) *BarbershopEmployeeRepository {
	return &BarbershopEmployeeRepository{db: db}
}

// Obtener los datos de un empleado
func (b *BarbershopEmployeeRepository) GetByID(id uint) (*models.BarbershopEmployee, error) {
	var barbershopEmployee models.BarbershopEmployee
	if err :=
		b.db.Preload("User").
			First(&barbershopEmployee, id).Error; err != nil {
		return nil, err
	}
	return &barbershopEmployee, nil
}
