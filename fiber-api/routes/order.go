package routes

import (
	"errors"
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

func CreateResponseOrder(order models.Order, user UserSerializer, product Product) Order {
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
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserReference)
		database.Database.Db.Find(&product, "id = ?", order.ProductReference)

		responseOrder := CreateResponseOrder(order, CreateUserResponse(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)

	}

	return c.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order not found")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order
	if err != nil {
		return c.Status(400).JSON("Cannot parse id, please provide a valid id")
	}
	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserReference)
	database.Database.Db.First(&product, order.ProductReference)

	responseUser := CreateUserResponse(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}
