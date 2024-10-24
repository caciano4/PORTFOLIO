package users

import (
	"github.com/caciano/portfolio/internal/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(db *gorm.DB, id string) (models.User, error) {
	var user models.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
