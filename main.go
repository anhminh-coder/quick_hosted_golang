package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"tiktok_info/database"
	"tiktok_info/entity"
)

func CreateUser(ctx *fiber.Ctx) error {
	userRequest := new(entity.UserRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	var user entity.User
	user.Name = userRequest.Name
	user.Phone = userRequest.Phone

	err := database.DB.Create(&user).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

func main() {
	database.DatabaseInit()
	database.RunMigration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001", // Adjust this based on your frontend address
		AllowMethods:     "GET,POST,HEAD,PUT,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Post("/users", CreateUser)
	app.Listen(":8888")
}
