package controllers

import (
	"GOlangRakamin/helpers"
	"GOlangRakamin/models"
	"GOlangRakamin/modules/auth"
	"GOlangRakamin/modules/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.Service
	authService auth.Service
}

func NewUserController(service user.Service, authService auth.Service) *userController {
	return &userController{service, authService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		//* menambil dari helpers
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Register Account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helpers.APIResponse("Register Account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helpers.APIResponse("Register Account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helpers.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {

	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)

	if err != nil {
		response := helpers.APIResponse("Login Account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helpers.APIResponse("Succesfully Loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}


func (h *userController) CheckEmailAvailability(c *gin.Context) {

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		//* menambil dari helpers
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helpers.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if IsEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helpers.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UploadAvatar(c *gin.Context) {
	// input dari user
	// simpan gambarnya di folder "images/"
	// di service kita panggil repo
	// JWT (sementara hardoce, seakan2 user yang login ID = 1)
	// repo ambil data user yang ID = 1
	// repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//* mendapatkan ID
	currentUser := c.MustGet("currentUser").(models.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helpers.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userController) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)

	formatter := user.FormatUser(currentUser, "")

	response := helpers.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
