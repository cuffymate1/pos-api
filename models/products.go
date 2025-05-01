package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Cost        float32  `json:"cost"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryId"`
}
