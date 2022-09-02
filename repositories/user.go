package repositories

import (
	"_waysbeans_/models"

	"gorm.io/gorm"
)

type UserRepository interface {
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, ID).Error

	return user, err
}
