package controller

import (
	"strconv"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/models"
	"github.com/cuffymate1/pos-api/services"
	"github.com/gofiber/fiber/v2"
)

func ListToppings(c *fiber.Ctx) error {
	toppings, err := services.ListToppings(config.GetDb())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(toppings)
}

func GetTopping(c *fiber.Ctx) error {
	var topping *models.Topping
	toppingId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	topping, err = services.GetTopping(config.GetDb(), uint(toppingId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(topping)
}

func CreateTopping(c *fiber.Ctx) error {
	topping := new(models.Topping)
	if err := c.BodyParser(&topping); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createTopping, err := services.CreateTopping(config.GetDb(), topping)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": createTopping,
	})
}

func UpdateTopping(c *fiber.Ctx) error {
	toppingId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	existingTopping := new(models.Topping)
	existingTopping, err = services.GetTopping(config.GetDb(), uint(toppingId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&existingTopping); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updateTopping, err := services.UpdateTopping(config.GetDb(), existingTopping)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": updateTopping,
	})
}

func DeleteTopping(c *fiber.Ctx) error {
	toppingId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	deleteTopping, err := services.DeleteTopping(config.GetDb(), uint(toppingId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": deleteTopping,
	})
}
