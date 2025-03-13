package services

import (
	"barberia/internal/models"
	"barberia/internal/repository"
	"errors"
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

// Obtener datos de todos los empleados registrados en la base de datos
func (b *BarbershopEmployeeService) GetAllEmployees() ([]models.BarbershopEmployee, int, error) {
	employees, err := b.BarbershopEmployeeRepo.GetAll()
	if err != nil {
		return nil, 0, err
	}
	return employees, len(employees), nil
}

// Borrar un empleado de la base de datos
func (b *BarbershopEmployeeService) DeleteEmployee(id uint) error {
	return b.BarbershopEmployeeRepo.Delete(id)
}

// Crear un nuevo empleado
func (b *BarbershopEmployeeService) CreateEmployee(employee *models.BarbershopEmployee) error {
	if employee.Barbershop_ID == 0 || employee.User_ID == 0 || employee.Comision_Porcentaje_Default < 0 || employee.Base_Salary < 0 {
		return errors.New("some data within the fields is null or zero")
	}
	return b.BarbershopEmployeeRepo.Create(employee)
}
