package topup

type SomethingWithID struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTopUp struct {
	UserID 		int 	`json:"user_id"`
	OrderID 	string	`json:"order_id"`
	GrossAmount int 	`json:"gross_amount"`
}