package category

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Category, error) {
	var category = []Category{}
	err := r.db.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
