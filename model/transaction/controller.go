package transaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoexsmass/helper"
)

type transactionController struct {
	transactionService Service
}

func NewTransactionController(transactionService Service) *transactionController {
	return &transactionController{transactionService}
}

func (h *transactionController) CreateTransaction(c *gin.Context) {
	var input CreateTransaction
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := TransactionFormatter(transaction)
	response := helper.APIResponse("Create Transaction Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) GetByUserID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get User Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetByUserID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get User Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := TransactionsFormatter(transactions)
	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) GetByID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transaction, err := h.transactionService.GetByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := TransactionFormatter(transaction)
	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}