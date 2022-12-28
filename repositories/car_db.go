package repositories

import (
	"car/models"
	"car/orms"

	"gorm.io/gorm"
)

type CarRepositoryDB struct {
	db *gorm.DB
}

// Create implements Car
func (r *CarRepositoryDB) Create(c *models.Car) (*models.Car, error) {
	orm := orms.ToCarORM(c)

	tx := r.db.Create(orm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orms.ToCar(orm), nil
}

// Delete implements Car
func (r *CarRepositoryDB) Delete(id int) error {
	orm := orms.Car{
		ID: id,
	}

	tx := r.db.Delete(orm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// FindByID implements Car
func (r *CarRepositoryDB) FindByID(id int) (*models.Car, error) {
	orm := &orms.Car{}

	tx := r.db.Find(orm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orms.ToCar(orm), nil
}

// FindMany implements Car
func (r *CarRepositoryDB) FindMany() ([]*models.Car, error) {
	orm := []*orms.Car{}

	tx := r.db.Find(&orm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orms.ToCars(orm), nil
}

// Update implements Car
func (r *CarRepositoryDB) Update(c *models.Car) (*models.Car, error) {
	orm := orms.ToCarORM(c)

	tx := r.db.Save(orm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orms.ToCar(orm), nil
}

func NewCarRepositoryDB(db *gorm.DB) Car {
	return &CarRepositoryDB{
		db: db,
	}
}
