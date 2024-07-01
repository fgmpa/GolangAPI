package repository

import (
	"GolangAPI"
	"github.com/jmoiron/sqlx"
)

type Car interface {
	Create(car GolangAPI.Car) (int, error)
	GetAll() ([]GolangAPI.Car, error)
	GetCar(id int) (GolangAPI.Car, error)
	Delete(id int) error
	Update(id int, input GolangAPI.UpdateCar) error
}
type Repository struct {
	Car
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Car: NewCarPostgres(db),
	}
}
