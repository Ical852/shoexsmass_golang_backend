package topup

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoexsmass/helper"
)

type topUpController struct {
	topUpService Service
}

func NewTopUpController(topUpService Service) *topUpController {
	return &topUpController{topUpService}
}

func (h *topUpController) TopUp(c *gin.Context) {
	var input CreateTopUp
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	topUp, err := h.topUpService.CreateTopUp(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatTopUp(topUp)

	response := helper.APIResponse("Create Top Up Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *topUpController) GetByUserID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get User Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	topUps, err := h.topUpService.GetByUserID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get User Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatTopUps(topUps)
	response := helper.APIResponse("Get User Top Up Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *topUpController) GetByID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	topUp, err := h.topUpService.GetByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatTopUp(topUp)
	response := helper.APIResponse("Get Top Up Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}