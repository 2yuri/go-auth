package user

import (
	"go-auth/database"
	"go-auth/models"
)

type UserRepository struct {
}

func (s UserRepository) Get(id uint) (models.User, error) {
	db := database.GetDatabase()

	var user models.User
	err := db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
func (s UserRepository) Create(user models.User) (models.User, error) {
	db := database.GetDatabase()
	err := db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
func (s UserRepository) Edit(user models.User) error {
	db := database.GetDatabase()

	err := db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}
func (s UserRepository) Delete(id uint) error {
	db := database.GetDatabase()
	err := db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
