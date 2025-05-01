package routes

import (
	"github.com/cuffymate1/pos-api/controller"
	"github.com/cuffymate1/pos-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {
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
}
