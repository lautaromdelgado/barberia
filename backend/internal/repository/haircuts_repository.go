package repository

import "gorm.io/gorm"

type HaircutsRepository struct {
	db *gorm.DB
}

func NewHaircutsRepository(db *gorm.DB) *HaircutsRepository {
	return &HaircutsRepository{db: db}
}
