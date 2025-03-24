package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nombre    string    `json:"nombre"`
	Apellido  string    `json:"apellido"`
	DNI       string    `json:"dni"`
	Correo    string    `json:"correo"`
	Rol       string    `json:"rol" gorm:"type:enum('owner','employee');default:'owner'"`
	Verified  bool      `json:"verified" gorm:"type:enum(1,0);default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "users"
}
