package service

import (
	"github.com/JavaHutt/rest-dealership/internal/model"
	"github.com/JavaHutt/rest-dealership/internal/repository"
)

type Service struct {
	Vehicle
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Vehicle: NewVehicleService(rep.Vehicle),
	}
}

type Vehicle interface {
	Get()
	GetAll() []model.Vehicle
	Create(model.Vehicle) (*model.Vehicle, error)
	Update()
	Delete()
}
