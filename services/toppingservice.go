package services

import (
	"github.com/cuffymate1/pos-api/models"
	"gorm.io/gorm"
)

func ListToppings(db *gorm.DB) ([]models.Topping, error) {
	var listToppings []models.Topping
	result := db.Find(&listToppings)
	if result.Error != nil {
		return nil, result.Error
	}

	return listToppings, nil
}

func GetTopping(db *gorm.DB, id uint) (*models.Topping, error) {
	var topping models.Topping
	result := db.First(&topping, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &topping, nil
}

func CreateTopping(db *gorm.DB, topping *models.Topping) (string, error) {
	result := db.Create(topping)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Created Successful", nil
}

func UpdateTopping(db *gorm.DB, topping *models.Topping) (string, error) {
	result := db.Save(topping)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Updated Successful", nil
}

func DeleteTopping(db *gorm.DB, id uint) (string, error) {
	var topping models.Topping
	if err := db.First(&topping, id); err != nil {
		return "Invalid Topping Id", err.Error
	}

	result := db.Delete(&topping, id)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Deleted Successful", nil
}
