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

// GetByID retorna un corte por su ID
func (h *HaircutsRepository) GetByID(id uint) (*models.Haircut, error) {
	var haircut models.Haircut
	if err := h.db.Preload("Barbershop").Preload("Barbershop.Owner").Preload("User").First(&haircut, id).Error; err != nil {
		return nil, err
	}
	return &haircut, nil
}

// Eliminar un corte
func (h *HaircutsRepository) Delete(id uint) error {
	if err := h.db.Delete(&models.Haircut{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Editar un corte
func (h *HaircutsRepository) Update(haircut *models.Haircut, id uint) error {
	if err := h.db.Model(&models.Haircut{}).Where("id = ?", id).Updates(haircut).Error; err != nil {
		return err
	}
	return nil
}

// Crear un corte
func (h *HaircutsRepository) Create(haircut *models.Haircut) error {
	if err := h.db.Create(haircut).Error; err != nil {
		return err
	}
	return nil
}
