package controller

import (
	"strconv"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/models"
	"github.com/cuffymate1/pos-api/services"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	users, err := services.ListUser(config.GetDb())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(users)
}

func Get(c *fiber.Ctx) error {
	var user *models.Users
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err = services.GetUser(config.GetDb(), uint(userId))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}

func Create(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createUser, err := services.CreateUser(config.GetDb(), user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": createUser,
	})
}

func Update(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	existingUser := new(models.Users)
	existingUser, err = services.GetUser(config.GetDb(), uint(userId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&existingUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updateUser, err := services.UpdateUser(config.GetDb(), existingUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": updateUser,
	})
}

func Delete(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	deleteUser, err := services.DeleteUser(config.GetDb(), uint(userId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": deleteUser,
	})
}
