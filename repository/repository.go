package repository

import (
	"github.com/Greeshmanth-Pulicallu/login-api/db"
	"github.com/Greeshmanth-Pulicallu/login-api/models"
)

func GetUserFromDB(userID string) (models.User, error) {
	var user models.User
	result := db.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, result.Error
}
