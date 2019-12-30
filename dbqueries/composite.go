package dbqueries

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"txCancel/models/db"
)

//go:generate mockgen -destination=mocks/mocks.go txCancel/dbqueries CompositeQI,CounterQI,TransactionQI,UserQI

// CompositeQI represents interface to perform db transactions
type CompositeQI interface {
	ApplyTransaction(transaction *db.Transaction) error
	CancelTransactions(txCount int) error
}

type CompositeQ struct {
	db *sqlx.DB

	userQ    UserQI
	txQ      TransactionQI
	counterQ CounterQI
}

func NewCompositeQ(db *sqlx.DB, userQ UserQI, txQ TransactionQI, counterQ CounterQI) *CompositeQ {
	return &CompositeQ{
		db:       db,
		userQ:    userQ,
		txQ:      txQ,
		counterQ: counterQ,
	}
}

func (q *CompositeQ) ApplyTransaction(transaction *db.Transaction) (err error) {
	tx, err := q.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback
			dErr := tx.Rollback()
			if dErr != nil {
				logrus.WithError(dErr).Error("failed to rollback")
			}
			err = errors.Errorf("panic: %v", p)
		} else if err != nil {
			// something went wrong, rollback
			dErr := tx.Rollback()
			if dErr != nil {
				logrus.WithError(dErr).Error("failed to rollback")
			}
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	amount := transaction.Amount
	// to decrease balance
	if transaction.State == db.StateLost {
		amount *= -1
	}

	err = q.userQ.UpdateBalanceTx(tx, 1, amount)
	if err != nil {
		return errors.Wrap(err, "failed to UpdateBalanceTx")
	}

	err = q.txQ.InsertTx(tx, transaction)
	if err != nil {
		return errors.Wrap(err, "failed to InsertTx")
	}

	return err
}

func (q *CompositeQ) CancelTransactions(txCount int) (err error) {
	tx, err := q.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback
			dErr := tx.Rollback()
			if dErr != nil {
				logrus.WithError(dErr).Error("failed to rollback")
			}
			err = errors.Errorf("panic: %v", p)
		} else if err != nil {
			// something went wrong, rollback
			dErr := tx.Rollback()
			if dErr != nil {
				logrus.WithError(dErr).Error("failed to rollback")
			}
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	// get last cancelled tx id
	counter, err := q.counterQ.GetByName(tx, db.CounterLastCancelledID)
	if err != nil {
		return errors.Wrap(err, "failed to get counter")
	}

	// get all txs after last cancelled
	txs, err := q.txQ.GetListGreaterThan(tx, counter.Value)
	if err != nil {
		return errors.Wrap(err, "failed to GetGreaterThan")
	}

	// first to cancel tx position (last)
	positionToCancel := len(txs) - 1

	// if count of new txs is odd and we have previously processed something (!= 0)
	// or if if count of new txs is even and we haven't previously processed something (== 0)
	// then second to last tx is the first to cancel
	if (len(txs)%2 != 0 && counter.Value != 0) || (len(txs)%2 == 0 && counter.Value == 0) {
		positionToCancel--
	}

	// cancel at most "txCount" txs
	for processed := 0; positionToCancel >= 0 && processed < txCount; positionToCancel -= 2 {
		t := txs[positionToCancel]
		amount := t.Amount
		// to decrease balance
		if t.State == db.StateWin {
			amount *= -1
		}

		// update balance
		err = q.userQ.UpdateBalanceTx(tx, 1, amount)
		if err != nil {
			return errors.Wrap(err, "failed to UpdateBalanceTx")
		}

		// update tx state
		err = q.txQ.CancelTx(tx, t.ID)
		if err != nil {
			return errors.Wrap(err, "failed to CancelTx")
		}

		// update last cancelled counter
		// condition because first iteration takes the most recent tx in a table
		if processed == 0 {
			err = q.counterQ.UpdateValueTx(tx, db.CounterLastCancelledID, t.ID)
			if err != nil {
				return errors.Wrap(err, "failed to update counter")
			}
		}

		processed++
	}

	return err
}
