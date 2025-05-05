package routes

import (
	"time"

	"github.com/cuffymate1/pos-api/controller"
	"github.com/cuffymate1/pos-api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func GetRoutes(app *fiber.App) {
	// ใช้ rate limiter กับทุก route
	app.Use(limiter.New(limiter.Config{
		Max:        5,                // ขอได้สูงสุด 5 ครั้ง
		Expiration: 30 * time.Second, // ภายใน 30 วินาที
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please try again later.",
			})
		},
	}))
	app.Post("login", controller.Login)
	app.Use(middleware.VerifyToken())
	// แยก Group
	users := app.Group("/User", middleware.OnlyAdmin())
	users.Get("/List", controller.List)
	users.Get("/:id", controller.Get)
	users.Post("/Create", controller.Create)
	users.Post("/Update/:id", controller.Update)
	users.Post("/Delete/:id", controller.Delete)

	products := app.Group("/Product")
	products.Get("/List", controller.ListProducts)
	products.Get("/:id", controller.GetProduct)
	products.Post("/Create", controller.CreateProduct)
	products.Post("/Update/:id", controller.UpdateProduct)
	products.Post("/Delete/:id", controller.DeleteProduct)

	categories := app.Group("/Category")
	categories.Get("/List", controller.ListCategory)
	categories.Get("/:id", controller.GetCategory)
	categories.Post("/Create", controller.CreateCategory)
	categories.Post("/Update/:id", controller.UpdateCategory)
	categories.Post("/Delete/:id", controller.DeleteCategory)

	toppings := app.Group("/Topping")
	toppings.Get("/List", controller.ListToppings)
	toppings.Get("/:id", controller.GetTopping)
	toppings.Post("/Create", controller.CreateTopping)
	toppings.Post("/Update/:id", controller.UpdateTopping)
	toppings.Post("/Delete/:id", controller.DeleteTopping)

	order := app.Group("/Order")
	order.Get("/List", controller.ListOrders)
	order.Get("/:id", controller.GetOrder)
	order.Post("/Create", controller.CreateOrder)
	order.Post("/Update/:id", controller.UpdateOrder)
	order.Post("/Delete/:id", controller.DeleteOrder)
}
