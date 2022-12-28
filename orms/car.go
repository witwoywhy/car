package orms

import "car/models"

type Car struct {
	ID       int `gorm:"primaryKey"`
	Brand    string
	Name     string
	Model    string
	SubModel string
	Color    string
	Price    int
}

func ToCarORM(c *models.Car) *Car {
	return &Car{
		ID:       c.ID,
		Brand:    c.Brand,
		Name:     c.Name,
		Model:    c.Model,
		SubModel: c.SubModel,
		Color:    c.Color,
		Price:    c.Price,
	}
}

func ToCar(orm *Car) *models.Car {
	return &models.Car{
		ID:       orm.ID,
		Brand:    orm.Brand,
		Name:     orm.Name,
		Model:    orm.Model,
		SubModel: orm.SubModel,
		Color:    orm.Color,
		Price:    orm.Price,
	}
}

func ToCars(orms []*Car) []*models.Car {
	cars := make([]*models.Car, len(orms))
	for i, orm := range orms {
		cars[i] = ToCar(orm)
	}
	return cars
}
