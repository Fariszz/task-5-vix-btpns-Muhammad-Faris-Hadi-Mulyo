package controllers

import (
	"GOlangRakamin/helpers"
	"GOlangRakamin/models"
	"GOlangRakamin/modules/photo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoService photo.Service
}

func NewPhotoController(photoService photo.Service) *photoController {
	return &photoController{photoService}
}

func (h *photoController) GetPhotos(c *gin.Context) {
	userID, _ := strconv.Atoi("user_id")
	photos, err := h.photoService.GetPhotos(userID)

	if err != nil {
		response := helpers.APIResponse("Failed to get photos", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("List of photos", http.StatusOK, "success", photo.FormatPhotos(photos))
	c.JSON(http.StatusOK, response)
}

func (h *photoController) GetPhoto(c *gin.Context) {
	var input photo.GetPhotoDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helpers.APIResponse("Failed to get detail of photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	photoDetail, err := h.photoService.GetPhotoByID(input)

	if err != nil {
		response := helpers.APIResponse("Failed to get detail of photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Detail of photo", http.StatusOK, "success", photo.FormatPhoto(photoDetail))
	c.JSON(http.StatusOK, response)
}

func (h *photoController) CreatePhoto(c *gin.Context) {
	var input photo.CreatePhotoInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Failed to create photo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUserr").(models.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	newPhoto, err := h.photoService.CreatePhotos(input, path)

	if err != nil {
		response := helpers.APIResponse("Failed to create photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success to create photo", http.StatusOK, "success", photo.FormatPhoto(newPhoto))
	c.JSON(http.StatusOK, response)
}

func (h *photoController) UpdatePhoto(c *gin.Context) {
	var inputID photo.GetPhotoDetailInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("Failed to update photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData photo.CreatePhotoInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Failed to update photo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedPhoto, err := h.photoService.UpdatePhoto(inputID, inputData)

	if err != nil {
		response := helpers.APIResponse("Failed to update photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success to update photo", http.StatusOK, "success", photo.FormatPhoto(updatedPhoto))
	c.JSON(http.StatusOK, response)
}

func (h *photoController) DeletePhoto(c *gin.Context) {
	var input photo.GetPhotoDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helpers.APIResponse("Failed to delete photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var userData photo.GetUserId

	currentUser := c.MustGet("currentUser").(models.User)
	userData.User = currentUser

	deletedPhoto, err := h.photoService.DeletePhotos(input, userData)

	if err != nil {
		response := helpers.APIResponse("Failed to delete photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success to delete photo", http.StatusOK, "success", deletedPhoto)
	c.JSON(http.StatusOK, response)
}
