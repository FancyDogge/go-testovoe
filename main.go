package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/foo", handleGet)
	app.Listen(":8000")
}

func handleGet(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "up and running"})
}
