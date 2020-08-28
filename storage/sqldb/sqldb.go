package sqldb

import (
	"context"
	"database/sql"
)

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).ExecContext() for additional documentation.
func Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("not implemented")
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).QueryContext() for additional documentation.
func Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	panic("not implemented")
}

// QueryRow executes a query that is expected to return at most one row.
//
// See (*database/sql.DB).QueryRowContext() for additional documentation.
func QueryRow(ctx context.Context, query string, args ...interface{}) *Row {
	panic("not implemented")
}

// Tx is a handle to a database transaction.
//
// See *database/sql.Tx for additional documentation.
type Tx struct{}

// Begin opens a new database transaction.
//
// See (*database/sql.DB).Begin() for additional documentation.
func Begin(ctx context.Context) (*Tx, error) {
	panic("not implemented")
}

// Commit commits the given transaction.
//
// See (*database/sql.Tx).Commit() for additional documentation.
func Commit(ctx context.Context, tx *Tx) error {
	panic("not implemented")
}

// Rollback rolls back the given transaction.
//
// See (*database/sql.Tx).Rollback() for additional documentation.
func Rollback(ctx context.Context, tx *Tx) error {
	panic("not implemented")
}

// ExecTx is like Exec but executes the query in the given transaction.
//
// See (*database/sql.Tx).ExecContext() for additional documentation.
func ExecTx(ctx context.Context, tx *Tx, query string, args ...interface{}) (sql.Result, error) {
	panic("not implemented")
}

// QueryTx is like Query but executes the query in the given transaction.
//
// See (*database/sql.Tx).QueryContext() for additional documentation.
func QueryTx(ctx context.Context, tx *Tx, query string, args ...interface{}) (*Rows, error) {
	panic("not implemented")
}

// QueryRowTx is like QueryRow but executes the query in the given transaction.
//
// See (*database/sql.Tx).QueryRowContext() for additional documentation.
func QueryRowTx(ctx context.Context, tx *Tx, query string, args ...interface{}) *Row {
	panic("not implemented")
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
	panic("not implemented")
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
//
// See (*database/sql.Rows).Scan() for additional documentation.
func (*Rows) Scan(dest ...interface{}) error {
	panic("not implemented")
}

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit Close.
//
// See (*database/sql.Rows).Err() for additional documentation.
func (*Rows) Err() error {
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
}
