package services

import (
	"os"
	"time"

	"github.com/cuffymate1/pos-api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB, user *models.Users) (string, error) {
	_ = godotenv.Load("config/.env")

	var existingUser models.Users
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return "Username is Invalid", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		return "Password is Invalid", err
	}

	t, err := createTokenJwt(&existingUser)
	if err != nil {
		return "Something went wrong!", err
	}

	return t, nil
}

func createTokenJwt(user *models.Users) (string, error) {
	JwtSecret := os.Getenv("JWT_SECRET")
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["fullname"] = user.Fullname
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "Something went wrong!", err
	}
	return t, nil
}
