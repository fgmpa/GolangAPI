package repository

import (
	"GolangAPI"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type CarPostgres struct {
	db *sqlx.DB
}

func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}

func (r *CarPostgres) Create(car GolangAPI.Car) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCarQuery := fmt.Sprintf("INSERT INTO %s (id,brand,model,ownersNum) VALUES ($1, $2, $3,$4) RETURNING id", carTable)
	row := tx.QueryRow(createCarQuery, car.Id, car.Brand, car.Model, car.OwnersNum)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *CarPostgres) GetAll() ([]GolangAPI.Car, error) {
	var cars []GolangAPI.Car
	query := fmt.Sprintf("SELECT * FROM %s", carTable)
	err := r.db.Select(&cars, query)
	return cars, err
}

func (r *CarPostgres) GetCar(id int) (GolangAPI.Car, error) {
	var car GolangAPI.Car
	query := fmt.Sprintf("SELECT * FROM %s ct WHERE ct.id = $1", carTable)
	err := r.db.Get(&car, query, id)
	return car, err
}
func (r *CarPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s ct WHERE ct.id = $1", carTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CarPostgres) Update(id int, input GolangAPI.UpdateCar) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Brand != nil {
		setValues = append(setValues, fmt.Sprintf("brand=$%d", argId))
		args = append(args, *input.Brand)
		argId++
	}
	if input.Model != nil {
		setValues = append(setValues, fmt.Sprintf("model=$%d", argId))
		args = append(args, *input.Model)
		argId++
	}
	if input.OwnersNum != nil {
		setValues = append(setValues, fmt.Sprintf("ownersNum=$%d", argId))
		args = append(args, *input.OwnersNum)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ct SET %s WHERE ct.id = $%d", carTable, setQuery, argId)
	args = append(args, id)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)
	_, err := r.db.Exec(query, args...)
	return err
}
