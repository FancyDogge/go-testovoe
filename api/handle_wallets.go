package api

import (
	"fmt"

	"github.com/FancyDogge/go-testovoe/db"
	"github.com/FancyDogge/go-testovoe/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WalletHandler struct {
	DB db.Database
}

// Создаем новый кошелек
func (h *WalletHandler) CreateWallet(c *fiber.Ctx) error {
	wallet := types.Wallet{}
	err := c.BodyParser(&wallet)
	fmt.Println(&wallet)
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	wallet.ID = uuid.New().String()
	wallet.Balance = 100.0

	err = h.DB.CreateWallet(wallet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.JSON(wallet)
}

// Транзакция
func (h *WalletHandler) TransferFunds(c *fiber.Ctx) error {
	walletID := c.Params("walletId")
	if _, err := h.DB.GetWallet(walletID); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Outgoing wallet not found"})
	}

	var request types.Transaction
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	err := h.DB.TransferFunds(walletID, request.To, request.Amount)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.SendStatus(200)
}

// История транзакций
func (h *WalletHandler) GetTransactionHistory(c *fiber.Ctx) error {
	walletID := c.Params("walletId")

	history, err := h.DB.GetTransactionHistory(walletID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Wallet not found"})
	}

	return c.JSON(history)
}

// Получаем кошель
func (h *WalletHandler) GetWallet(c *fiber.Ctx) error {
	walletID := c.Params("walletId")

	wallet, err := h.DB.GetWallet(walletID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Wallet not found"})
	}

	return c.JSON(wallet)
}
