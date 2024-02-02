package main

import (
	"flag"
	"log"

	"github.com/FancyDogge/go-testovoe/api"
	"github.com/FancyDogge/go-testovoe/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "The listen port of the API server") // --listenAddr ":7777" -> будет работать на порту :7777
	flag.Parse()

	// !!! Я извиняюсь за такой бред без .env, но я зашиваюсь из-за работы и не успеваю к дедлайну, сдам прямо на последней минуте
	connectionUri := "mongodb://admin:adminpassword@localhost:27017"
	dbName := "ewallet"
	collectionName := "wallets"

	mongoDB, err := db.NewMongoDB(connectionUri, dbName, collectionName)
	if err != nil {
		log.Fatal(err)
	}

	walletHandler := &api.WalletHandler{
		DB: mongoDB,
	}

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Post("/wallet", walletHandler.CreateWallet)
	apiv1.Post("/wallet/:walletId/send", walletHandler.TransferFunds)
	apiv1.Get("/wallet/:walletId/history", walletHandler.GetTransactionHistory)
	apiv1.Get("/wallet/:walletId", walletHandler.GetWallet)

	if app.Listen(*listenAddr); err != nil {
		panic(err)
	}
}
