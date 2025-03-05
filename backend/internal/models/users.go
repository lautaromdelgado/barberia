package models

import "time"

type User struct {
	ID        uint      `json:"id_user" gorm:"primaryKey"`
	Nombre    string    `json:"nombre"`
	Apellido  string    `json:"apellido"`
	DNI       string    `json:"dni"`
	Correo    *string   `json:"correo"`
	Rol       string    `json:"rol" gorm:"type:enum('owner','employee')"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (u User) TableName() string {
	return "users"
}
