package models

import "time"

type Barbershop struct {
	ID        uint      `json:"id_barbershop" gorm:"primaryKey"`
	Owner_ID  uint      `json:"id_owner"`
	Nombre    string    `json:"nombre"`
	Direccion string    `json:"direccion"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	Owner User `json:"owner" gorm:"foreignKey:Owner_ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (b Barbershop) TableName() string {
	return "barbershops"
}
