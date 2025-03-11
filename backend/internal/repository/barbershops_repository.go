package repository

import (
	"barberia/internal/models"

	"gorm.io/gorm"
)

// Barbershops es un repositorio para la tabla barbershops
type Barbershops struct {
	db *gorm.DB
}

// NewBarbershopsRepository crea un nuevo repositorio para la tabla barbershops
func NewBarbershopsRepository(db *gorm.DB) *Barbershops {
	return &Barbershops{db: db}
}

// GetAll retorna todos los barbershops de la base de datos
func (b *Barbershops) GetAll() ([]models.Barbershop, error) {
	var barbershops []models.Barbershop
	if err := b.db.Find(&barbershops).Error; err != nil {
		return nil, err
	}
	return barbershops, nil
}
