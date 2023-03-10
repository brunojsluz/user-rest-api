package controller

import (
	"github.com/gofiber/fiber/v2"
	"user-rest-api/database"
	"user-rest-api/models"
)

func FindAll(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func FindByCode(c *fiber.Ctx) error {
	var users []models.User
	var code = c.Params("code")
	database.DB.First(&users, code)
	return c.JSON(users)
}

func CreateUSer(c *fiber.Ctx) error {
	var user = new(models.User)
	c.BodyParser(user)
	database.DB.Create(&user)
	return c.JSON(user)
}
