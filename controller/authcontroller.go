package controller

import (
	"time"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/models"
	"github.com/cuffymate1/pos-api/services"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginUser := new(models.LoginModel)
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := &models.Users{
		Username:     loginUser.Username,
		PasswordHash: loginUser.Password,
	}

	token, err := services.Auth(config.GetDb(), user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": token,
		})
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Jwt",
		Value:    token,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,
		Secure:   true,   // ใช้ true ถ้าใช้ https
		SameSite: "None", // หรือ "Strict" / "None" ตามความต้องการ
	})

	return c.JSON(fiber.Map{
		"message": "Login Succesful",
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "Jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	return c.JSON(fiber.Map{
		"message": "Logout Succesful",
	})
}
