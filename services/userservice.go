package services

import (
	"regexp"
	"strings"

	"github.com/cuffymate1/pos-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// PasswordStrengthError represents password validation errors
type PasswordStrengthError struct {
	Message string
}

func (e *PasswordStrengthError) Error() string {
	return e.Message
}

// validatePasswordStrength checks if the password meets the required strength criteria
func validatePasswordStrength(password string) error {
	var errors []string

	if len(password) < 8 {
		errors = append(errors, "รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		errors = append(errors, "รหัสผ่านต้องมีตัวอักษรพิมพ์ใหญ่อย่างน้อย 1 ตัว")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		errors = append(errors, "รหัสผ่านต้องมีตัวอักษรพิมพ์เล็กอย่างน้อย 1 ตัว")
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		errors = append(errors, "รหัสผ่านต้องมีตัวเลขอย่างน้อย 1 ตัว")
	}

	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		errors = append(errors, "รหัสผ่านต้องมีตัวอักษรพิเศษอย่างน้อย 1 ตัว")
	}

	if len(errors) > 0 {
		return &PasswordStrengthError{
			Message: strings.Join(errors, "\n"),
		}
	}

	return nil
}

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
	// Validate password strength
	if err := validatePasswordStrength(user.PasswordHash); err != nil {
		return "Password need 8 character, 1 uppercase, 1 lowercase, 1 number, 1 special character", err
	}

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
		// Validate password strength
		if err := validatePasswordStrength(user.PasswordHash); err != nil {
			return "Password need 8 character, 1 uppercase, 1 lowercase, 1 number, 1 special character", err
		}

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
