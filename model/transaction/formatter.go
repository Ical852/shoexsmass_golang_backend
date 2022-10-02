package transaction

import (
	"shoexsmass/model/product"
	"shoexsmass/model/transactiondetail"
	"shoexsmass/model/user"
)

type TransactionDetailFormat struct {
	ID 				int							`json:"id"`
	TransactionID 	int							`json:"transaction_id"`
	ProductID 		int							`json:"product_id"`
	Qty 			int							`json:"qty"`
	Product 		product.ProductFormatter	`json:"product"`
}

func TDetailFormatter(tDetail transactiondetail.TransactionDetail) TransactionDetailFormat {
	tdDetailFormatted := TransactionDetailFormat{
		ID:            tDetail.ID,
		TransactionID: tDetail.TransactionID,
		ProductID:     tDetail.ProductID,
		Qty:           tDetail.Qty,
		Product:       product.FormatProduct(tDetail.Product),
	}

	return tdDetailFormatted
}

func TDetailsFormatter(tDetails []transactiondetail.TransactionDetail) []TransactionDetailFormat {
	var tdDetailsFormatted []TransactionDetailFormat
	for _, tDetail := range tDetails {
		tdDetailsFormatted = append(tdDetailsFormatted, TDetailFormatter(tDetail))
	}

	return tdDetailsFormatted
}

type TransactionFormat struct {
	ID 					int							`json:"id"`
	UserID 				int							`json:"user_id"`
	PayMethod 			string						`json:"pay_method"`
	OrderID 			string						`json:"order_id"`
	GrossAmount 		int							`json:"gross_amount"`
	Status 				string						`json:"status"`
	PaymentUrl 			string						`json:"payment_url"`
	User 				user.UserFormatter			`json:"user"`
	TransactionDetail 	[]TransactionDetailFormat	`json:"transaction_detail"`
}

func TransactionFormatter(transaction Transaction) TransactionFormat {
	transactionFormatted := TransactionFormat{
		ID:                transaction.ID,
		UserID:            transaction.UserID,
		PayMethod:         transaction.PayMethod,
		OrderID:           transaction.OrderID,
		GrossAmount:       transaction.GrossAmount,
		Status:            transaction.Status,
		PaymentUrl:        transaction.PaymentUrl,
		User:              user.FormatUser(transaction.User),
		TransactionDetail: TDetailsFormatter(transaction.TransactionDetail),
	}

	return transactionFormatted
}

func TransactionsFormatter(transactions []Transaction) []TransactionFormat {
	var transactionsFormatted []TransactionFormat
	for _, transaction := range transactions {
		transactionsFormatted = append(transactionsFormatted, TransactionFormatter(transaction))
	}

	return transactionsFormatted
}
