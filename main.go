package main

import (
	"flag"

	"github.com/FancyDogge/go-testovoe/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "The listen port of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleGetFoo)
	apiv1.Get("/wallet/:id", api.HandleGetWallet)
	app.Listen(*listenAddr) // --listenAddr ":7777" -> будет работать на порту :7777
}

func handleGetFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "up and running Foo"})
}
