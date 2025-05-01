package controller

import (
	"strconv"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/models"
	"github.com/cuffymate1/pos-api/services"
	"github.com/gofiber/fiber/v2"
)

func ListProducts(c *fiber.Ctx) error {
	product, err := services.ListProducts(config.GetDb())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product *models.Product
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messsage": err.Error(),
		})
	}

	product, err = services.GetProduct(config.GetDb(), uint(productId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createProduct, err := services.CreateProduct(config.GetDb(), product)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": createProduct,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	existingProduct := new(models.Product)
	existingProduct, err = services.GetProduct(config.GetDb(), uint(productId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&existingProduct); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updateProduct, err := services.UpdateProduct(config.GetDb(), existingProduct)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": updateProduct,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	deltedProduct, err := services.DeleteProduct(config.GetDb(), uint(productId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": deltedProduct,
	})
}
