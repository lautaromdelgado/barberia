package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
)

// HaircutsServices es el servicio de cortes
type HaircutsServices struct {
	HaircutsRepository *repository.HaircutsRepository
}

// Constructor de HaircutsServices
func NewHaircutsServices(haircutsRepository *repository.HaircutsRepository) *HaircutsServices {
	return &HaircutsServices{
		HaircutsRepository: haircutsRepository,
	}
}

// GetAllHaircuts retorna todos los cortes
func (h *HaircutsServices) GetAllHaircuts() ([]models.Haircut, error) {
	return h.HaircutsRepository.GetAll()
}

// GetHaircutByID retorna un corte por su ID
func (h *HaircutsServices) GetHaircutByID(id uint) (*models.Haircut, error) {
	return h.HaircutsRepository.GetByID(id)
}
