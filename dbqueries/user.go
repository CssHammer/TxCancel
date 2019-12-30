package dbqueries

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"

	"txCancel/models/db"
)

const TableUsers = "users"

// UserQI represents interface to access users
type UserQI interface {
	GetByID(id int) (*db.User, error)
	UpdateBalanceTx(tx *sqlx.Tx, id int, amount float64) error // used in DB transaction
}

type UserQ struct {
	db       *sqlx.DB
	sBuilder sq.SelectBuilder
}

func NewUserQ(db *sqlx.DB) *UserQ {
	return &UserQ{
		db:       db,
		sBuilder: sq.Select("*").From(TableUsers),
	}
}

func (q *UserQ) GetByID(id int) (*db.User, error) {
	var user db.User

	query, args, err := q.sBuilder.Where("id = ?", id).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sql")
	}

	query = q.db.Rebind(query)
	err = q.db.Get(&user, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from db")
	}

	return &user, nil
}

func (q *UserQ) UpdateBalanceTx(tx *sqlx.Tx, id int, amount float64) error {

	// using plain SQL due to squirrel issues with increment
	query := fmt.Sprintf("UPDATE %v SET balance = balance + ? WHERE id = ?", TableUsers)
	args := []interface{}{amount, id}

	query = tx.Rebind(query)
	_, err := tx.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec UpdateBalanceTx")
	}

	return nil
}
