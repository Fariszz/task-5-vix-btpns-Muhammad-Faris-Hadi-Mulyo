package photo

import (
	"GOlangRakamin/database"
	"GOlangRakamin/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Photo, error)
	FindByUserID(userID int) ([]models.Photo, error)
	FindByID(ID int) (models.Photo, error)
	Save(product models.Photo) (models.Photo, error)
	Update(product models.Photo) (models.Photo, error)
	Delete(product models.Photo) (models.Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{
		db: database.DbManager(),
	}
}

func (r *repository) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil

}

func (r *repository) FindByUserID(userID int) ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil

}

func (r *repository) FindByID(ID int) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Preload("User").Where("id = ?", ID).Find(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil

}

func (r *repository) Save(photo models.Photo) (models.Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil

}

func (r *repository) Update(photo models.Photo) (models.Photo, error) {
	err := r.db.Save(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil

}

func (r *repository) Delete(photo models.Photo) (models.Photo, error) {
	err := r.db.Delete(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil

}
