package repositories

import "car/models"

type Car interface {
	Create(*models.Car) (*models.Car, error)
	FindByID(int) (*models.Car, error)
	FindMany() ([]*models.Car, error)
	Update(*models.Car) (*models.Car, error)
	Delete(int) error
}
