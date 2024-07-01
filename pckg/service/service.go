package service

import (
	"GolangAPI"
	"GolangAPI/pckg/repository"
)

type Car interface {
	Create(car GolangAPI.Car) (int, error)
	GetAll() ([]GolangAPI.Car, error)
	GetCar(id int) (GolangAPI.Car, error)
	Delete(id int) error
	Update(id int, input GolangAPI.UpdateCar) error
}
type Service struct {
	Car
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Car: NewCarService(repos.Car),
	}
}
