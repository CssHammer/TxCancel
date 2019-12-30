package db

type User struct {
	ID      int64   `db:"id"`
	Name    string  `db:"name"`
	Balance float64 `db:"balance"`
}
