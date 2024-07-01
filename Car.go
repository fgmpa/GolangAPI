package GolangAPI

import "errors"

type Car struct {
	Id        int    `json:"id" db:"id"`
	Brand     string `json:"brand" db:"brand"`
	Model     string `json:"model" db:"model"`
	OwnersNum int    `json:"ownersNum" json:"ownersNum"`
}

type UpdateCar struct {
	Brand     *string `json:"brand"`
	Model     *string `json:"model"`
	OwnersNum *int    `json:"ownersNum"`
}

func (i UpdateCar) Validate() error {
	if i.Model == nil && i.Brand == nil && i.OwnersNum == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
