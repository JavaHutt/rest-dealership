package repository

import (
	"github.com/JavaHutt/rest-dealership/internal/model"

	"gorm.io/gorm"
)

type VehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func (rep *VehicleRepository) Get(id int) (*model.Vehicle, error) {
	vehicle := new(model.Vehicle)

	result := rep.db.First(vehicle, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return vehicle, nil
}

func (rep *VehicleRepository) GetAll() []model.Vehicle {
	var vehicles []model.Vehicle
	rep.db.Find(&vehicles)

	return vehicles
}

func (rep *VehicleRepository) Create(vehicle model.Vehicle) (*model.Vehicle, error) {
	result := rep.db.Create(&vehicle)

	if result.Error != nil {
		return nil, result.Error
	}
	return &vehicle, nil
}
func (rep *VehicleRepository) Update() {

}
func (rep *VehicleRepository) Delete(id int) error {
	return rep.db.Delete(&model.Vehicle{}, id).Error
}
