package repository

import (
	"barberia/internal/models"

	"gorm.io/gorm"
)

// HaircutsRepository es el repositorio de cortes
type HaircutsRepository struct {
	db *gorm.DB
}

// Constructor de HaircutsRepository
func NewHaircutsRepository(db *gorm.DB) *HaircutsRepository {
	return &HaircutsRepository{db: db}
}

// GetAll retorna todos los cortes
func (h *HaircutsRepository) GetAll() ([]models.Haircut, error) {
	var haircuts []models.Haircut
	if err := h.db.Preload("Barbershop").Preload("Barbershop.Owner").Preload("User").Find(&haircuts).Error; err != nil {
		return nil, err
	}
	return haircuts, nil
}
