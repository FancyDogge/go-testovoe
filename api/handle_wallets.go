package api

import (
	"github.com/FancyDogge/go-testovoe/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetWallet(c *fiber.Ctx) error {
	w := types.Wallet{
		Balance: 1234.0,
	}

	return c.JSON(w)
}
