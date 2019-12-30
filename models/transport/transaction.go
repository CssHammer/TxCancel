package transport

type Transaction struct {
	State         string `db:"state" json:"state"`
	Amount        string `db:"amount" json:"amount"`
	TransactionID string `db:"transaction_id" json:"transaction_id"`
}
