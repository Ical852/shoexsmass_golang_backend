package transaction

import (
	"shoexsmass/model/transactiondetail"
	"shoexsmass/model/user"
)

type Transaction struct {
	ID int
	UserID int
	PayMethod string
	OrderID string
	GrossAmount int
	Status string
	PaymentUrl string
	User user.User
	TransactionDetail []transactiondetail.TransactionDetail
}