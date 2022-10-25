package user

import "GOlangRakamin/models"

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageUrl string `json:"image_url"`
}

func FormatUser(user models.User, token string) UserFormatter {

	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
		ImageUrl: user.Avatar,
	}

	return formatter
}
