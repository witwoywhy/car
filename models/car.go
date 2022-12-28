package models

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var errBrandInvalid = errors.New("brand invalid")
var errorNameInvalid = errors.New("name invalid")
var errorModelInvalid = errors.New("model invalid")
var errorColorInvalid = errors.New("color invalid")

type Car struct {
	ID       int
	Brand    string
	Name     string
	Model    string
	SubModel string
	Color    string
	Price    int
}

func NewCar(brand, name, model, submodel, color string, price int) (*Car, error) {
	if utf8.RuneCountInString(brand) == 0 {
		return nil, errBrandInvalid
	}

	if utf8.RuneCountInString(name) == 0 {
		return nil, errorNameInvalid
	}

	if utf8.RuneCountInString(model) == 0 {
		return nil, errorModelInvalid
	}

	if utf8.RuneCountInString(color) == 0 {
		return nil, errorColorInvalid
	}

	return &Car{
		Brand:    strings.ToLower(brand),
		Name:     strings.ToLower(name),
		Model:    strings.ToLower(model),
		SubModel: strings.ToLower(submodel),
		Color:    strings.ToLower(color),
		Price:    price,
	}, nil
}
