package product

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Product, error)
	GetByCategoryID(categoryID int) ([]Product, error)
	GetByID(ID int) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Preload("ProductImage").Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetByCategoryID(categoryID int) ([]Product, error) {
	var products []Product
	err := r.db.Where("category_id = ?", categoryID).Preload("ProductImage").Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetByID(ID int) (Product, error) {
	var product Product
	err := r.db.Where("id = ?", ID).Preload("ProductImage").Preload("Category").Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
