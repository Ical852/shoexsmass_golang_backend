package transactiondetail

import "gorm.io/gorm"

type Repository interface {
	Create(transactionDetail TransactionDetail) (TransactionDetail, error)
	GetByTransactionID(transactionID int) ([]TransactionDetail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(transactionDetail TransactionDetail) (TransactionDetail, error) {
	err := r.db.Create(&transactionDetail).Error
	if err != nil {
		return transactionDetail, err
	}

	return transactionDetail, nil
}

func (r *repository) GetByTransactionID(transactionID int) ([]TransactionDetail, error) {
	var transactionDetails []TransactionDetail
	err := r.db.Preload("Product").Find(&transactionDetails).Error
	if err != nil {
		return transactionDetails, err
	}

	return transactionDetails, nil
}
