package photo

import "GOlangRakamin/models"

type Service interface {
	GetPhotos(userID int) ([]models.Photo, error)
	GetPhotoByID(input GetPhotoDetailInput) (models.Photo, error)
	CreatePhotos(input CreatePhotoInput, filelocation string) (models.Photo, error)
	UpdatePhoto(input GetPhotoDetailInput, inputData CreatePhotoInput) (models.Photo, error)
	DeletePhotos(input GetPhotoDetailInput, userData GetUserId) (models.Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPhotos(userID int) ([]models.Photo, error) {
	if userID != 0 {
		photos, err := s.repository.FindByUserID(userID)
		if err != nil {
			return photos, err
		}
		return photos, nil
	}

	photos, err := s.repository.FindAll()
	if err != nil {
		return photos, err
	}
	return photos, nil
}

func (s *service) GetPhotoByID(input GetPhotoDetailInput) (models.Photo, error) {
	photo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *service) CreatePhotos(input CreatePhotoInput, filelocation string) (models.Photo, error) {
	photo := models.Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = filelocation
	photo.UserId = input.UserId

	newPhoto, err := s.repository.Save(photo)

	if err != nil {
		return newPhoto, err
	}
	return newPhoto, nil
}

func (s *service) UpdatePhoto(input GetPhotoDetailInput, inputData CreatePhotoInput) (models.Photo, error) {
	photo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return photo, err
	}

	photo.Title = inputData.Title
	photo.PhotoUrl = inputData.PhotoUrl
	photo.Caption = inputData.Caption

	updatedPhoto, err := s.repository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}
	return updatedPhoto, nil
}

func (s *service) DeletePhotos(input GetPhotoDetailInput, userData GetUserId) (models.Photo, error) {
	photo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return photo, err
	}

	if photo.UserId != userData.User.ID {
		return photo, nil
	}

	deletedPhoto, err := s.repository.Delete(photo)
	if err != nil {
		return deletedPhoto, err
	}
	return deletedPhoto, nil
}