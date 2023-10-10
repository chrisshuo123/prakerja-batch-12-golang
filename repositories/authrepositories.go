package repositories

import (
	"prakerja12/configs"
	"prakerja12/models"
)

func Register(user *models.User) error {
	result := configs.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}