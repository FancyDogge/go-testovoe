package types

type Wallet struct {
	ID      string  `bson:"_id" json:"id"`
	Balance float64 `bson:"balance" json:"balance"`
}
