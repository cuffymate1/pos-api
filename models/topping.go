package models

import (
	"gorm.io/gorm"
)

type Topping struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
