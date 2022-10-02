package transactiondetail

import "shoexsmass/model/product"

type TransactionDetail struct {
	ID int
	TransactionID int
	ProductID int
	Qty int
	Product product.Product
}