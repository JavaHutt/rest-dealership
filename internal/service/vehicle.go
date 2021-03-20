package service

import (
	"github.com/JavaHutt/rest-dealership/internal/model"
	"github.com/JavaHutt/rest-dealership/internal/repository"
)

type VehicleService struct {
	rep repository.Vehicle
}

func NewVehicleService(rep repository.Vehicle) *VehicleService {
	return &VehicleService{
		rep: rep,
	}
}

func (s *VehicleService) Get(id int) (*model.Vehicle, error) {
	return s.rep.Get(id)
}

func (s *VehicleService) GetAll() []model.Vehicle {
	return s.rep.GetAll()
}

func (s *VehicleService) Create(vehicle model.Vehicle) (*model.Vehicle, error) {
	return s.rep.Create(vehicle)
}

func (s *VehicleService) Update(id int, vehicle model.Vehicle) (*model.Vehicle, error) {
	return s.rep.Update(id, vehicle)
}

func (s *VehicleService) Delete(id int) error {
	return s.rep.Delete(id)
}

func (s *VehicleService) SeedData(vehicles []model.Vehicle) error {
	for _, vehicle := range vehicles {
		if _, err := s.rep.Create(vehicle); err != nil {
			return err
		}
	}
	return nil
}
