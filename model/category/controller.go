package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoexsmass/helper"
)

type catagoryController struct {
	catagoryService Service
}

func NewCategoryController(userService Service) *catagoryController {
	return &catagoryController{userService}
}

func (h *catagoryController) GetALl(c *gin.Context) {
	categories, err := h.catagoryService.GetALl()
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Categories Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatCategories(categories)

	response := helper.APIResponse("Get Categories Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}