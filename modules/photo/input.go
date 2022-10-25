package photo

import (
	"GOlangRakamin/models"
)

type GetPhotoDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetUserId struct {
	User models.User
}

type CreatePhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserId   int    `json:"user_id" binding:"required"`
	User     models.User
}
