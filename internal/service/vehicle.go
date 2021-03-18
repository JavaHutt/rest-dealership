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

func (s *VehicleService) Get() {

}
func (s *VehicleService) GetAll() {
	s.rep.GetAll()
}
func (s *VehicleService) Create(vehicle model.Vehicle) (*model.Vehicle, error) {
	return s.rep.Create(vehicle)
}
func (s *VehicleService) Update() {

}
func (s *VehicleService) Delete() {

}
