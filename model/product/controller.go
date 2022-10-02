package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoexsmass/helper"
)

type productController struct {
	productService Service
}

func NewProductController(productService Service) *productController {
	return &productController{productService}
}

func (h *productController) GetAll(c *gin.Context) {
	products, err := h.productService.GetAll()
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Products Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatProducts(products)

	response := helper.APIResponse("Get Products Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *productController) GetByCategoryID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Products By Category ID Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.productService.GetByCategoryID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Products By Category ID Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatProducts(products)

	response := helper.APIResponse("Get Products By Category ID Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *productController) GetByID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Products By ID Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.productService.GetByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Products By ID Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatProduct(product)

	response := helper.APIResponse("Get Products By ID Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}