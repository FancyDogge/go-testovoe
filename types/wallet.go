package types

import "time"

type Wallet struct {
	ID           string        `bson:"_id" json:"id"`
	Balance      float64       `bson:"balance" json:"balance"`
	Transactions []Transaction `json:"transactions" bson:"transactions"`
}

type Transaction struct {
	Time   time.Time `json:"time" bson:"time"`
	From   string    `json:"from" bson:"from"`
	To     string    `json:"to" bson:"to"`
	Amount float64   `json:"amount" bson:"amount"`
}
