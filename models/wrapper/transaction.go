package wrapper

import (
	"strconv"

	"github.com/pkg/errors"

	"txCancel/models/db"
	"txCancel/models/transport"
)

func UnwrapTx(tx *transport.Transaction) (*db.Transaction, error) {
	amount, err := strconv.ParseFloat(tx.Amount, 64)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse amount")
	}

	return &db.Transaction{
		State:         tx.State,
		Amount:        amount,
		TransactionID: tx.TransactionID,
		Canceled:      false,
	}, nil
}
