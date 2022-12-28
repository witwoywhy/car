package repositories

import (
	"car/models"
	"errors"
	"sync"
)

var errorCarNotFound = errors.New("car not found")

type CarRepositoryMock struct {
	a []*models.Car
	h map[int]*models.Car

	mu sync.Mutex
}

// Create implements Car
func (r *CarRepositoryMock) Create(c *models.Car) (*models.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	c.ID = len(r.a) + 1

	r.a = append(r.a, c)
	r.h[c.ID] = c

	return c, nil
}

// Delete implements Car
func (r *CarRepositoryMock) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.a = append(r.a[:id-1], r.a[id:]...)
	delete(r.h, id)

	return nil
}

// FindByID implements Car
func (r *CarRepositoryMock) FindByID(id int) (*models.Car, error) {
	if c, ok := r.h[id]; ok {
		return c, nil
	} else {
		return nil, errorCarNotFound
	}
}

// FindMany implements Car
func (r *CarRepositoryMock) FindMany() ([]*models.Car, error) {
	return r.a, nil
}

// Update implements Car
func (r *CarRepositoryMock) Update(c *models.Car) (*models.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.a[c.ID-1] = c
	r.h[c.ID] = c
	return c, nil
}

func NewCarRepositoryMock() Car {
	return &CarRepositoryMock{
		a:  []*models.Car{},
		h:  make(map[int]*models.Car),
		mu: sync.Mutex{},
	}
}
