package main

import (
	"log"

	"github.com/Esilahic/fiber-api/database"
	"github.com/Esilahic/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API")
}

func setupRoutes(app *fiber.App) {
	//welcome route
	app.Get("/api", welcome)
	//user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	//product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	//order endpoints

}

func main() {
	database.Connect()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
