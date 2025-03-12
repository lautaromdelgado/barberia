package models

import "time"

type BarbershopEmployee struct {
	ID                          uint      `json:"id_barbershop_employee" gorm:"primaryKey"`
	Barbershop_ID               uint      `json:"id_barbershop"`
	User_ID                     uint      `json:"id_user"`
	Comision_Porcentaje_Default float64   `json:"comision" gorm:"type:decimal(5,2);not null;default:0"`      // % de comisión
	Base_Salary                 float64   `json:"salario_base" gorm:"type:decimal(10,2);not null;default:0"` // Salario base
	Created_At                  time.Time `json:"created_at" gorm:"autoCreateTime"`

	Barbershop *Barbershop `json:"barbershop,omitempty" gorm:"foreignKey:Barbershop_ID;references:ID"`
	User       User        `json:"user" gorm:"foreignKey:User_ID;references:ID"`
}

func (b BarbershopEmployee) TableName() string {
	return "barbershop_employees"
}
