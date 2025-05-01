package services

import (
	"github.com/cuffymate1/pos-api/models"
	"gorm.io/gorm"
)

func ListCategory(db *gorm.DB) ([]models.Category, error) {
	var listCategories []models.Category
	result := db.Preload("Product").Find(&listCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	return listCategories, nil
}

func GetCategory(db *gorm.DB, id uint) (*models.Category, error) {
	var category models.Category
	result := db.Preload("Product").First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func CreateCategory(db *gorm.DB, category *models.Category) (string, error) {
	result := db.Create(category)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Created Successful", nil
}

func UpdateCategory(db *gorm.DB, category *models.Category) (string, error) {
	result := db.Save(category)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Updated Successful", nil
}

func DeleteCategory(db *gorm.DB, id uint) (string, error) {
	var category models.Category
	if err := db.First(&category, id); err != nil {
		return "Invalid Category Id", err.Error
	}

	result := db.Delete(&category, id)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Deleted Successful", nil
}
