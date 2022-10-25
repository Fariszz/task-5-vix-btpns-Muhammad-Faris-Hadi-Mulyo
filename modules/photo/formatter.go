package photo

import "GOlangRakamin/models"

type PhotoFormatter struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"string"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

func FormatPhoto(photos models.Photo) PhotoFormatter {
	photoFormatter := PhotoFormatter{}
	photoFormatter.ID = photos.ID
	photoFormatter.Title = photos.Title
	photoFormatter.Caption = photos.Caption
	photoFormatter.PhotoUrl = photos.PhotoUrl
	photoFormatter.UserId = photos.UserId

	return photoFormatter

}

func FormatPhotos(photos []models.Photo) []PhotoFormatter {
	var photosFormatter []PhotoFormatter

	for _, photo := range photos {
		photoFormatter := FormatPhoto(photo)
		photosFormatter = append(photosFormatter, photoFormatter)
	}

	return photosFormatter
}
