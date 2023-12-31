package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"path/filepath"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func Test(ctx *fiber.Ctx) error {
	var user User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "bad request",
		})
	}

	filePath := filepath.Join(".", "data.txt")

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "opening file error",
		})
	}
	defer file.Close()
	_, err = file.WriteString("name: " + user.Name + ", phone: " + user.Phone + "\n")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "writing file error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "successfully",
	})
}

func ServeFile(ctx *fiber.Ctx) error {
	filePath := filepath.Join(".", "data.txt")

	return ctx.SendFile(filePath)
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // Adjust this based on your frontend address
		AllowMethods:     "GET,POST,HEAD,PUT,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/file", ServeFile)
	app.Post("/test", Test)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Home")
	})
	app.Listen(":8080")
}
