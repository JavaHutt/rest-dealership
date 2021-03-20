package repository

import (
	"github.com/JavaHutt/rest-dealership/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	Vehicle
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Vehicle: NewVehicleRepository(db),
	}
}

type Vehicle interface {
	Get(id int) (*model.Vehicle, error)
	GetAll() []model.Vehicle
	Create(vehicle model.Vehicle) (*model.Vehicle, error)
	Update(id int, vehicle model.Vehicle) (*model.Vehicle, error)
	Delete(id int) error
}
