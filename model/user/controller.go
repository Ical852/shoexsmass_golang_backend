package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shoexsmass/helper"
)

type userController struct {
	userService Service
}

func NewUserController(userService Service) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := FormatUser(newUser)

	response := helper.APIResponse("Register Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context)  {
	var input LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUser(loggedinUser)

	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) CheckEmailAvailability(c *gin.Context) {
	var input CheckEmail

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UploadAvatar(c *gin.Context) {
	var input UploadAvatar
	err := c.ShouldBindUri(&input)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := input.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded" : true}
	response := helper.APIResponse("Avatar Success Uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userController) FetchUser(c *gin.Context) {
	var input UploadAvatar
	err := c.ShouldBindUri(&input)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Get User", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.GetUserByID(input.ID)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Get User", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := FormatUser(user)

	response := helper.APIResponse("Successfuly Fetch User Data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
	var input FormUpdateUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Update User", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userUpdated, err := h.userService.UpdateUser(input)
	if err != nil {
		data := gin.H{"is_uploaded" : false, "error" : err.Error()}
		response := helper.APIResponse("Failed to Update User", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUser(userUpdated)
	response := helper.APIResponse("Successfuly Updated User Data", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}