package notification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoexsmass/helper"
)

type notificationController struct {
	notificationService Service
}

func NewNotificationController(notificationService Service) *notificationController {
	return &notificationController{notificationService}
}

func (h *notificationController) CreateNotif(c *gin.Context) {
	var input NotificationCreateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Notif Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	notif, err := h.notificationService.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Notif Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatNotification(notif)
	response := helper.APIResponse("Create Notif Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *notificationController) GetByID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Notif Detail Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	notif, err := h.notificationService.GetByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Notif Detail Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatNotification(notif)
	response := helper.APIResponse("Get Notif Detail Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *notificationController) Delete(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Delete Notif Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	notif, err := h.notificationService.Delete(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Delete Notif Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatNotification(notif)
	response := helper.APIResponse("Delete Notif Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *notificationController) GetByUserID(c *gin.Context) {
	var input SomethingWithID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get User Notif Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	notifications, err := h.notificationService.GetByUserID(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get User Notif Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatNotifications(notifications)
	response := helper.APIResponse("Get User Notif Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}