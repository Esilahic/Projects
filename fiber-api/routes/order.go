package routes

import (
	"time"

	"github.com/Esilahic/fiber-api/database"
	"github.com/Esilahic/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint           `json:"id"`
	User      UserSerializer `json:"user"`
	Product   Product        `json:"product"`
	CreatedAt time.Time      `json:"created_at"`
}

func OrderResponse(order models.Order, user UserSerializer, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(order.UserReference, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductReference, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)
	responseUser := CreateUserResponse(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := OrderResponse(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}
