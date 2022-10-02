package transaction

type SomethingWithID struct {
	ID int `uri:"id" binding:"required"`
}

type TDDetailInput struct {
	ProductID int 	`json:"product_id"`
	Qty int 		`json:"qty"`
}

type CreateTransaction struct {
	UserID 			int 			`json:"user_id"`
	OrderID 		string			`json:"order_id"`
	GrossAmount 	int 			`json:"gross_amount"`
	PayMethod 		string 			`json:"pay_method"`
	TdDetail		[]TDDetailInput `json:"td_detail"`
}