package topup

import "shoexsmass/model/user"

type TopUp struct {
	ID int
	UserID int
	OrderID string
	GrossAmount int
	Status string
	PaymentUrl string
	User user.User
}