package product

import (
	"shoexsmass/model/category"
	"shoexsmass/model/product_image"
)

type Product struct {
	ID int
	Name string
	CategoryID int
	Desc string
	Price int
	Category category.Category
	ProductImage []product_image.ProductImage
}