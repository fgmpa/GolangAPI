package service

import (
	"GolangAPI"
	"GolangAPI/pckg/repository"
)

type CarService struct {
	repo repository.Car
}

func NewCarService(repo repository.Car) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) Create(car GolangAPI.Car) (int, error) {
	return s.repo.Create(car)
}

func (s *CarService) GetAll() ([]GolangAPI.Car, error) {
	return s.repo.GetAll()
}

func (s *CarService) GetCar(id int) (GolangAPI.Car, error) {
	return s.repo.GetCar(id)
}
func (s *CarService) Delete(id int) error {
	return s.repo.Delete(id)
}
func (s *CarService) Update(id int, input GolangAPI.UpdateCar) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
