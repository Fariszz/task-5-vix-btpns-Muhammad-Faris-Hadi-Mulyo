package user

import (
	"GOlangRakamin/database"
	"GOlangRakamin/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
	FindByID(ID int) (models.User, error)
	Update(user models.User)(models.User, error)
	FindByEmail(email string) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{
		db: database.DbManager(),
	}
}

func (r * repository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r * repository) FindByID(ID int) (models.User, error){
	var user models.User
	err := r.db.Where("ID = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}