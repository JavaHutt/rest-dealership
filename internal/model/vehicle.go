package model

type VehicleStatus int8

const (
	StatusInTransit VehicleStatus = iota
	StatusInStock
	StatusSold
	StatusWithdrawn
)

type Vehicle struct {
	Brand   string        `json:"brand"   db:"brand"`
	Model   string        `json:"model"   db:"model"`
	Price   uint          `json:"price"   db:"price"`
	Status  VehicleStatus `json:"status"  db:"status"`
	Mileage uint          `json:"mileage" db:"mileage"`
}

func (vs VehicleStatus) String() string {
	return [4]string{"В пути", "На складе", "Продан", "Снят с продажи"}[vs]
}
