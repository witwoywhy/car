package services

import (
	"car/models"
	"car/repositories"
)

type CarService struct {
	repo repositories.Car
}

func (s *CarService) Create(brand, name, model, subModel, color string, price int) (*models.Car, error) {
	c, err := models.NewCar(brand, name, model, subModel, color, price)
	if err != nil {
		return nil, err
	}

	c, err = s.repo.Create(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CarService) FindByID(id int) (*models.Car, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CarService) FindMany() ([]*models.Car, error) {
	c, err := s.repo.FindMany()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CarService) Update(id int, brand, name, model, subModel, color string, price int) (*models.Car, error) {
	c, err := models.NewCar(brand, name, model, subModel, color, price)
	if err != nil {
		return nil, err
	}
	c.ID = id

	c, err = s.repo.Update(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CarService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewCarService(repo repositories.Car) *CarService {
	return &CarService{
		repo: repo,
	}
}
