package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
	"errors"
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

// Eliminar un corte por su ID
func (h *HaircutsServices) DeleteHaircut(id uint) error {
	return h.HaircutsRepository.Delete(id)
}

// Actualizar un corte
func (h *HaircutsServices) UpdateHaircut(haircut *models.Haircut, id uint) error {
	if haircut.BarbershopID == 0 && haircut.UserID == 0 && haircut.MontoTotal < 0 && haircut.ComisionAplicada < 0 && haircut.PorcentajeComision < 0 {
		return errors.New("no se ha enviado ningún campo para actualizar")
	}
	return h.HaircutsRepository.Update(haircut, id)
}
