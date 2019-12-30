package dbqueries

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"

	"txCancel/models/db"
)

const TableCounters = "counters"

// CounterQI represents interface to access counters
type CounterQI interface {
	GetByName(tx *sqlx.Tx, name string) (*db.Counter, error) // used in DB transaction
	UpdateValueTx(tx *sqlx.Tx, name string, value int) error // used in DB transaction
}

type CounterQ struct {
	db       *sqlx.DB
	sBuilder sq.SelectBuilder
	uBuilder sq.UpdateBuilder
}

func NewCounterQ(db *sqlx.DB) *CounterQ {
	return &CounterQ{
		db:       db,
		sBuilder: sq.Select("*").From(TableCounters),
		uBuilder: sq.Update(TableCounters),
	}
}

func (q *CounterQ) GetByName(tx *sqlx.Tx, name string) (*db.Counter, error) {
	var counter db.Counter

	query, args, err := q.sBuilder.Where("name = ?", name).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sql")
	}

	query = tx.Rebind(query)
	err = tx.Get(&counter, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get counter from db")
	}

	return &counter, nil
}

func (q *CounterQ) UpdateValueTx(tx *sqlx.Tx, name string, value int) error {

	query, args, err := q.uBuilder.Set("value", value).Where("name = ?", name).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to get sql")
	}

	query = tx.Rebind(query)
	_, err = tx.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec UpdateValueTx")
	}

	return nil
}
