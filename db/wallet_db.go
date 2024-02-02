package db

import (
	"context"
	"time"

	"github.com/FancyDogge/go-testovoe/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Интерфейс без хардкода определенной бд
type Database interface {
	CreateWallet(wallet types.Wallet) error
	TransferFunds(from, to string, amount float64) error
	GetTransactionHistory(walletID string) ([]types.Transaction, error)
	GetWallet(walletID string) (types.Wallet, error)
}

type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoDB(uri, dbName, collectionName string) (*MongoDB, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &MongoDB{
		client:     client,
		database:   db,
		collection: collection,
	}, nil
}

func (m *MongoDB) CreateWallet(wallet types.Wallet) error {
	if wallet.Transactions == nil {
		wallet.Transactions = []types.Transaction{}
	}

	_, err := m.collection.InsertOne(context.Background(), wallet)
	return err
}

func (m *MongoDB) TransferFunds(from, to string, amount float64) error {
	// Обновляем сендера
	_, err := m.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": from},
		bson.M{"$inc": bson.M{"balance": -amount}},
	)
	if err != nil {
		return err
	}

	// Обновляем получателя
	_, err = m.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": to},
		bson.M{"$inc": bson.M{"balance": amount}},
	)
	if err != nil {
		return err
	}

	// Записываем транзакцию сендера
	transactionFrom := types.Transaction{
		Time:   time.Now(),
		From:   from,
		To:     to,
		Amount: -amount,
	}
	_, err = m.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": from},
		bson.M{"$push": bson.M{"transactions": transactionFrom}},
	)
	if err != nil {
		return err
	}

	// Записываем транзакцию получателя
	transactionTo := types.Transaction{
		Time:   time.Now(),
		From:   from,
		To:     to,
		Amount: amount,
	}
	_, err = m.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": to},
		bson.M{"$push": bson.M{"transactions": transactionTo}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) GetTransactionHistory(walletID string) ([]types.Transaction, error) {
	var wallet types.Wallet
	err := m.collection.FindOne(context.Background(), bson.M{"_id": walletID}).Decode(&wallet)
	if err != nil {
		return nil, err
	}

	return wallet.Transactions, nil
}

func (m *MongoDB) GetWallet(walletID string) (types.Wallet, error) {
	var wallet types.Wallet
	err := m.collection.FindOne(context.Background(), bson.M{"_id": walletID}).Decode(&wallet)
	if err != nil {
		return types.Wallet{}, err
	}

	return wallet, nil
}
