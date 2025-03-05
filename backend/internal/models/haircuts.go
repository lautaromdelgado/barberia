package models

import "time"

type Haircut struct {
	ID                 uint      `json:"id_corte" gorm:"primaryKey"`
	BarbershopID       uint      `json:"id_barbershop"`
	UserID             uint      `json:"id_user"` // Puede ser dueño o empleado
	RealizadoEn        time.Time `json:"realizado_en" gorm:"not null;default:CURRENT_TIMESTAMP"`
	MontoTotal         float64   `json:"monto_total" gorm:"type:decimal(10,2);not null"`
	ComisionAplicada   float64   `json:"comision_aplicada" gorm:"type:decimal(10,2);not null;default:0"`  // Comisión si es empleado
	PorcentajeComision float64   `json:"porcentaje_comision" gorm:"type:decimal(5,2);not null;default:0"` // % Comisión guardado para histórico
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`

	Barbershop Barbershop `json:"barbershop" gorm:"foreignKey:BarbershopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User       User       `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (h Haircut) TableName() string {
	return "haircuts"
}
