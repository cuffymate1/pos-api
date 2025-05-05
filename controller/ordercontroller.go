package controller

import (
	"strconv"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/models"
	"github.com/cuffymate1/pos-api/services"
	"github.com/gofiber/fiber/v2"
)

func ListOrders(c *fiber.Ctx) error {
	orders, err := services.ListOrders(config.GetDb())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order, err := services.GetOrders(config.GetDb(), uint(orderId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(order)
}

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createOrder, err := services.CreateOrder(config.GetDb(), order)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": createOrder,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Get existing order
	existingOrder := new(models.Order)
	if err := config.GetDb().First(existingOrder, orderId).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Parse new order data
	if err := c.BodyParser(existingOrder); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updateOrder, err := services.UpdateOrder(config.GetDb(), existingOrder)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": updateOrder,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	deletedOrder, err := services.DeleteOrder(config.GetDb(), uint(orderId))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": deletedOrder,
	})
}
