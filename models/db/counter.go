package db

const (
	CounterLastCancelledID = "last_cancelled_id"
)

type Counter struct {
	Name  string `db:"name"`
	Value int    `db:"value"`
}
