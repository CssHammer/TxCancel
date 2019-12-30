package db

const (
	StateWin  = "win"
	StateLost = "lost"
)

type Transaction struct {
	ID            int     `db:"id" json:"id"`
	State         string  `db:"state" json:"state"`
	Amount        float64 `db:"amount" json:"amount"`
	TransactionID string  `db:"transaction_id" json:"transaction_id"`
	Canceled      bool    `db:"canceled" json:"canceled"`
}
