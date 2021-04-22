package sqldb

import (
	"context"
	"database/sql"
)

// An error satisfying ErrNoRows is reported by Scan
// when QueryRow doesn't return a row.
// It must be tested against with errors.Is.
var ErrNoRows = sql.ErrNoRows

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).ExecContext() for additional documentation.
func Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("encore apps must be run using the encore command")
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).QueryContext() for additional documentation.
func Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	panic("encore apps must be run using the encore command")
}

// QueryRow executes a query that is expected to return at most one row.
//
// See (*database/sql.DB).QueryRowContext() for additional documentation.
func QueryRow(ctx context.Context, query string, args ...interface{}) *Row {
	panic("encore apps must be run using the encore command")
}

// Tx is a handle to a database transaction.
//
// See *database/sql.Tx for additional documentation.
type Tx struct{}

// Begin opens a new database transaction.
//
// See (*database/sql.DB).Begin() for additional documentation.
func Begin(ctx context.Context) (*Tx, error) {
	panic("encore apps must be run using the encore command")
}

// Commit commits the given transaction.
//
// See (*database/sql.Tx).Commit() for additional documentation.
func Commit(tx *Tx) error {
	panic("encore apps must be run using the encore command")
}

// Rollback rolls back the given transaction.
//
// See (*database/sql.Tx).Rollback() for additional documentation.
func Rollback(tx *Tx) error {
	panic("encore apps must be run using the encore command")
}

// ExecTx is like Exec but executes the query in the given transaction.
//
// See (*database/sql.Tx).ExecContext() for additional documentation.
func ExecTx(tx *Tx, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("encore apps must be run using the encore command")
}

// QueryTx is like Query but executes the query in the given transaction.
//
// See (*database/sql.Tx).QueryContext() for additional documentation.
func QueryTx(tx *Tx, ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	panic("encore apps must be run using the encore command")
}

// QueryRowTx is like QueryRow but executes the query in the given transaction.
//
// See (*database/sql.Tx).QueryRowContext() for additional documentation.
func QueryRowTx(tx *Tx, ctx context.Context, query string, args ...interface{}) *Row {
	panic("encore apps must be run using the encore command")
}

// Rows is the result of a query. Its cursor starts before the first row
// of the result set. Use Next to advance from row to row.
//
// See *database/sql.Rows for additional documentation.
type Rows struct{}

// Close closes the Rows, preventing further enumeration.
//
// See (*database/sql.Rows).Close() for additional documentation.
func (*Rows) Close() error {
	panic("encore apps must be run using the encore command")
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
//
// See (*database/sql.Rows).Scan() for additional documentation.
func (*Rows) Scan(dest ...interface{}) error {
	panic("encore apps must be run using the encore command")
}

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit Close.
//
// See (*database/sql.Rows).Err() for additional documentation.
func (*Rows) Err() error {
	panic("encore apps must be run using the encore command")
}

// Next prepares the next result row for reading with the Scan method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. Err should be consulted to distinguish between
// the two cases.
//
// Every call to Scan, even the first one, must be preceded by a call to Next.
//
// See (*database/sql.Rows).Next() for additional documentation.
func (*Rows) Next() bool {
	panic("encore apps must be run using the encore command")
}

// Row is the result of calling QueryRow to select a single row.
//
// See *database/sql.Row for additional documentation.
type Row struct{}

// Scan copies the columns from the matched row into the values
// pointed at by dest.
//
// See (*database/sql.Row).Scan() for additional documentation.
func (*Row) Scan(dest ...interface{}) error {
	panic("encore apps must be run using the encore command")
}
