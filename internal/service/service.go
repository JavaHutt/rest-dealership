package service

import "github.com/JavaHutt/rest-dealership/internal/repository"

type Service struct {
	Vehicle
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}

type Vehicle interface{}
