package repositories

import (
	"_waysbeans_/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfiles() ([]models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
	CreateProfile(product models.Profile) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfiles() ([]models.Profile, error) {
	var products []models.Profile
	err := r.db.Preload("User").Find(&products).Error

	return products, err

}
func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var product models.Profile
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *repository) CreateProfile(product models.Profile) (models.Profile, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) UpdateProfile(product models.Profile) (models.Profile, error) {
	err := r.db.Save(&product).Error

	return product, err
}
