package models

import "time"

type Barbershop struct {
	ID        uint      `json:"id_barbershop" gorm:"primaryKey"`
	OwnerID   uint      `json:"id_owner" gorm:"column:owner_id"`
	Nombre    string    `json:"nombre"`
	Direccion string    `json:"direccion"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	Owner User `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
}

func (b Barbershop) TableName() string {
	return "barbershops"
}
