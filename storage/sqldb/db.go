package sqldb

import (
	"context"
	"database/sql"
)

type Database struct {
	_ int // for godoc to show unexported fields
}

// Stdlib returns a *sql.DB object that is connected to the same db,
// for use with libraries that expect a *sql.DB.
func (*Database) Stdlib() *sql.DB {
	panic("encore apps must be run using the encore command")
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).ExecContext() for additional documentation.
func (*Database) Exec(ctx context.Context, query string, args ...interface{}) (ExecResult, error) {
	panic("encore apps must be run using the encore command")
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).QueryContext() for additional documentation.
func (*Database) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	panic("encore apps must be run using the encore command")
}

// QueryRow executes a query that is expected to return at most one row.
//
// See (*database/sql.DB).QueryRowContext() for additional documentation.
func (*Database) QueryRow(ctx context.Context, query string, args ...interface{}) *Row {
	panic("encore apps must be run using the encore command")
}

// Begin opens a new database transaction.
//
// See (*database/sql.DB).Begin() for additional documentation.
func (*Database) Begin(ctx context.Context) (*Tx, error) {
	panic("encore apps must be run using the encore command")
}
