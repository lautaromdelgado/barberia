package database

import (
	"log"

	"barberia/internal/models"

	"gorm.io/gorm"
)

// RunMigrations ejecuta la migración de las tablas en la base de datos.
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Barbershop{},
		&models.BarbershopEmployee{},
		&models.Haircut{},
	)
	if err != nil {
		log.Fatalf("Error en la migración: %v", err)
	}
	log.Println("✅ Migraciones completadas exitosamente")
}
