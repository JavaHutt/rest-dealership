package repository

import "gorm.io/gorm"

type Repository struct {
	Vehicle
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}

type Vehicle interface{}
