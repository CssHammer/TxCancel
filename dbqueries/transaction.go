package dbqueries

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"

	"txCancel/models/db"
)

const TableTransactions = "transactions"

// TransactionQI represents interface to access game transactions
type TransactionQI interface {
	InsertTx(tx *sqlx.Tx, transaction *db.Transaction) error             // used in DB transaction
	GetListGreaterThan(tx *sqlx.Tx, limit int) ([]db.Transaction, error) // used in DB transaction
	CancelTx(tx *sqlx.Tx, id int) error                                  // used in DB transaction
}

type TransactionQ struct {
	db       *sqlx.DB
	uBuilder sq.UpdateBuilder
	iBuilder sq.InsertBuilder
}

func NewTransactionQ(db *sqlx.DB) *TransactionQ {
	return &TransactionQ{
		db:       db,
		uBuilder: sq.Update(TableTransactions),
		iBuilder: sq.Insert(TableTransactions),
	}
}

func (q *TransactionQ) InsertTx(tx *sqlx.Tx, transaction *db.Transaction) error {
	query, args, err := q.iBuilder.SetMap(map[string]interface{}{
		"state":          transaction.State,
		"amount":         transaction.Amount,
		"transaction_id": transaction.TransactionID,
		"canceled":       transaction.Canceled,
	}).ToSql()

	if err != nil {
		return errors.Wrap(err, "failed to get sql")
	}

	query = tx.Rebind(query)

	_, err = tx.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec InsertTx")
	}

	return nil
}

func (q *TransactionQ) GetListGreaterThan(tx *sqlx.Tx, startID int) ([]db.Transaction, error) {
	var txs []db.Transaction

	// using plain SQL due to squirrel issues with FOR UPDATE
	query := fmt.Sprintf("SELECT * FROM %v "+
		"WHERE id > ? "+
		"ORDER BY id "+
		"FOR UPDATE",
		TableTransactions)
	args := []interface{}{startID}

	query = tx.Rebind(query)
	err := tx.Select(&txs, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select txs for cancel")
	}

	return txs, nil
}

func (q *TransactionQ) CancelTx(tx *sqlx.Tx, id int) error {

	query, args, err := q.uBuilder.SetMap(map[string]interface{}{
		"canceled": true,
	}).Where("id = ?", id).ToSql()

	if err != nil {
		return errors.Wrap(err, "failed to get sql")
	}

	query = tx.Rebind(query)
	_, err = tx.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec CancelTx")
	}

	return nil
}
