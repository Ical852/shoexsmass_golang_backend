package product

import (
	"shoexsmass/model/category"
	"shoexsmass/model/product_image"
)

type ProductImageFormatter struct {
	ID 			int		`json:"id"`
	ProductID 	int		`json:"product_id"`
	Image 		string	`json:"image"`
}

type ProductFormatter struct {
	ID 				int							`json:"id"`
	Name 			string						`json:"name"`
	CategoryID 		int							`json:"category_id"`
	Desc 			string						`json:"desc"`
	Price 			int							`json:"price"`
	Category 		category.CategoryFormatter	`json:"category"`
	ProductImage 	[]ProductImageFormatter		`json:"product_image"`
}

func FormatProductImage(productImage product_image.ProductImage) ProductImageFormatter {
	productImageFormatted := ProductImageFormatter{
		ID:        productImage.ID,
		ProductID: productImage.ProductID,
		Image:     productImage.Image,
	}

	return productImageFormatted
}

func FormatProductImages(productImages []product_image.ProductImage) []ProductImageFormatter {
	var productImagesFormatted []ProductImageFormatter
	for _, image := range productImages {
		productImagesFormatted = append(productImagesFormatted, FormatProductImage(image))
	}

	return productImagesFormatted
}

func FormatProduct(product Product) ProductFormatter {
	productFormatted := ProductFormatter{
		ID:           product.ID,
		Name:         product.Name,
		CategoryID:   product.CategoryID,
		Desc:         product.Desc,
		Price:        product.Price,
		Category:     category.FormatCategory(product.Category),
		ProductImage: FormatProductImages(product.ProductImage),
	}

	return productFormatted
}

func FormatProducts(products []Product) []ProductFormatter {
	var productsFormatted []ProductFormatter
	for _, product := range products {
		productsFormatted = append(productsFormatted, FormatProduct(product))
	}

	return productsFormatted
}