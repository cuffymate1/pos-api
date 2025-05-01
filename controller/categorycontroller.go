package controller

import (
	"strconv"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/models"
	"github.com/cuffymate1/pos-api/services"
	"github.com/gofiber/fiber/v2"
)

func ListCategory(c *fiber.Ctx) error {
	categories, err := services.ListCategory(config.GetDb())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(categories)
}

func GetCategory(c *fiber.Ctx) error {
	var categories *models.Category
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	categories, err = services.GetCategory(config.GetDb(), uint(categoryId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createCategory, err := services.CreateCategory(config.GetDb(), category)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": createCategory,
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	existingCategory := new(models.Category)
	existingCategory, err = services.GetCategory(config.GetDb(), uint(categoryId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&existingCategory); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updateCategory, err := services.UpdateCategory(config.GetDb(), existingCategory)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": updateCategory,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	deleteCategory, err := services.DeleteCategory(config.GetDb(), uint(categoryId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": deleteCategory,
	})
}
