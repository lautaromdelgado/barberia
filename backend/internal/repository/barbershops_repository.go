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
	if err := b.db.Preload("Owner").Find(&barbershops).Error; err != nil {
		return nil, err
	}
	return barbershops, nil
}

// GetByID retorna un barbershop por su ID
func (b *Barbershops) GetByID(id uint) (*models.Barbershop, error) {
	var barbershop models.Barbershop
	if err := b.db.Preload("Owner").First(&barbershop, id).Error; err != nil {
		return nil, err
	}
	return &barbershop, nil
}

// Crear una nueva barbería
func (b *Barbershops) Create(barbershop *models.Barbershop) error {
	if err := b.db.Create(barbershop).Error; err != nil {
		return err
	}
	return nil
}
