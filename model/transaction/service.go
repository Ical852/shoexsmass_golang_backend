package transaction

import (
	"shoexsmass/model/transactiondetail"
)

type Service interface {
	CreateTransaction(input CreateTransaction) (Transaction, error)
	GetByUserID(input SomethingWithID) ([]Transaction, error)
	GetByID(input SomethingWithID) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input CreateTransaction) (Transaction, error) {
	transaction := Transaction{
		ID:                0,
		UserID:            input.UserID,
		PayMethod:         input.PayMethod,
		OrderID:           input.OrderID,
		GrossAmount:       input.GrossAmount,
	}



	var tdDetails []transactiondetail.TransactionDetail

	for _, detailInput := range input.TdDetail {
		tdDetails = append(tdDetails, transactiondetail.TransactionDetail{
			ProductID:     detailInput.ProductID,
			Qty:           detailInput.Qty,
		})
	}

	transaction.TransactionDetail = tdDetails
	transaction.OrderID = input.OrderID

	transactionCreate, err := s.repository.Create(transaction)
	if err != nil {
		return transactionCreate, err
	}

	return transactionCreate, nil
}

func (s *service) GetByUserID(input SomethingWithID) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetByID(input SomethingWithID) (Transaction, error) {
	transactions, err := s.repository.GetByID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
