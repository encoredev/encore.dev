package sqldb

import (
	"context"
	"database/sql"
)

func Begin(ctx context.Context) (*sql.Tx, error) {
	panic("not implemented")
}

func Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("not implemented")
}

func Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	panic("not implemented")
}

func QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	panic("not implemented")
}
