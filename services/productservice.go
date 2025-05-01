package services

import (
	"github.com/cuffymate1/pos-api/models"
	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) ([]models.Product, error) {
	var listProduct []models.Product
	result := db.Preload("Category").Find(&listProduct)
	if result.Error != nil {
		return nil, result.Error
	}

	return listProduct, nil
}

func GetProduct(db *gorm.DB, id uint) (*models.Product, error) {
	var product models.Product
	result := db.Preload("Category").First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func CreateProduct(db *gorm.DB, product *models.Product) (string, error) {
	var category models.Category
	if err := db.First(&category, product.CategoryId).Error; err != nil {
		return "Invalid category ID", err
	}

	result := db.Create(product)

	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Created Successful", nil
}

func UpdateProduct(db *gorm.DB, product *models.Product) (string, error) {
	var category models.Category
	if err := db.First(&category, product.CategoryId).Error; err != nil {
		return "Invalid category ID", err
	}

	result := db.Save(product)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Updated Successful", nil
}

func DeleteProduct(db *gorm.DB, id uint) (string, error) {
	var product models.Product
	if err := db.First(&product, id); err != nil {
		return "Invalid Product Id", err.Error
	}

	result := db.Delete(&product, id)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}

	return "Deleted Successful", nil
}
