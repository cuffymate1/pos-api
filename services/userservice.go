package services

import (
	"github.com/cuffymate1/pos-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) ([]models.Users, error) {
	var listUsers []models.Users
	result := db.Find(&listUsers)

	if result.Error != nil {
		return nil, result.Error
	}

	return listUsers, nil
}

func GetUser(db *gorm.DB, id uint) (*models.Users, error) {
	var getUser models.Users
	result := db.First(&getUser, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &getUser, nil
}

func CreateUser(db *gorm.DB, user *models.Users) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return "Something Went Wrong!", err
	}
	user.PasswordHash = string(hashedPass)
	result := db.Create(user)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Create Successful", nil
}

func UpdateUser(db *gorm.DB, user *models.Users) (string, error) {
	if user.PasswordHash != "" {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return "Failed to hash password", err
		}
		user.PasswordHash = string(hashedPass)
	}

	result := db.Save(user)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Update Successful", nil
}

func DeleteUser(db *gorm.DB, id uint) (string, error) {
	var user models.Users
	if err := db.First(&user, id); err != nil {
		return "Invalid User Id", err.Error
	}
	result := db.Delete(&user, id)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Delete Successful", nil
}
