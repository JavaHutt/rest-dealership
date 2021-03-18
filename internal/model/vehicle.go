package model

import (
	"errors"

	"gorm.io/gorm"
)

type VehicleStatus int8

const (
	StatusInTransit VehicleStatus = iota
	StatusInStock
	StatusSold
	StatusWithdrawn
)

type Vehicle struct {
	gorm.Model
	Brand        string        `json:"brand"           db:"brand"`
	VehicleModel string        `json:"vehicle_model"   db:"vehicle_model"`
	Price        uint          `json:"price"           db:"price"`
	Status       VehicleStatus `json:"status"          db:"status"        gorm:"default:0"`
	Mileage      uint          `json:"mileage"         db:"mileage"       gorm:"default:0"`
}

func (vs VehicleStatus) String() string {
	return [4]string{"В пути", "На складе", "Продан", "Снят с продажи"}[vs]
}

func (v *Vehicle) BeforeCreate(tx *gorm.DB) error {
	if v.Brand == "" || len(v.Brand) > 100 {
		return errors.New("invalid brand field")
	}
	if v.VehicleModel == "" || len(v.VehicleModel) > 100 {
		return errors.New("invalid model field")
	}
	if v.Price == 0 {
		return errors.New("missing price field, no free cars!")
	}
	return nil
}
