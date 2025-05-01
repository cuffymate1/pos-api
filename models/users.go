package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"passwordhash"`
	Fullname     string `json:"fullname"`
	Role         string `json:"role"`
}
