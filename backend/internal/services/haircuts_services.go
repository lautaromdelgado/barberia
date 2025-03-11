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
