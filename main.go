package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Test(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello")
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001", // Adjust this based on your frontend address
		AllowMethods:     "GET,POST,HEAD,PUT,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Post("/test", Test)
	app.Listen("0.0.0.0:8888")
}
