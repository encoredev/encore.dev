// Package sqldb provides Encore services direct access to their databases.
//
// For the documentation on how to use databases within Encore see https://encore.dev/docs/develop/databases.
package sqldb

import (
	"context"
	"database/sql"
)

// ErrNoRows is an error reported by Scan when QueryRow doesn't return a row.
// It must be tested against with errors.Is.
var ErrNoRows = sql.ErrNoRows

// ExecResult is the result of an Exec query.
type ExecResult interface {
	// RowsAffected returns the number of rows affected. If the result was not
	// for a row affecting command (e.g. "CREATE TABLE") then it returns 0.
	RowsAffected() int64
}

// Tx is a handle to a database transaction.
//
// See *database/sql.Tx for additional documentation.
type Tx struct {
	_ int // for godoc to show unexported fields
}

// Commit commits the given transaction.
//
// See (*database/sql.Tx).Commit() for additional documentation.
func (*Tx) Commit() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L41-L41
	doPanic("encore apps must be run using the encore command")
	return
}

// Rollback rolls back the given transaction.
//
// See (*database/sql.Tx).Rollback() for additional documentation.
func (*Tx) Rollback() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L46-L46
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Tx) Exec(ctx context.Context, query string, args ...interface{}) (_ ExecResult, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L84-L86
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Tx) Query(ctx context.Context, query string, args ...interface{}) (_ *Rows, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L113-L139
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Tx) QueryRow(ctx context.Context, query string, args ...interface{}) (_ *Row) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L141-L167
	doPanic("encore apps must be run using the encore command")
	return
}

// Rows is the result of a query. Its cursor starts before the first row
// of the result set. Use Next to advance from row to row.
//
// See *database/sql.Rows for additional documentation.
type Rows struct {
	_ int // for godoc to show unexported fields
}

// Close closes the Rows, preventing further enumeration.
//
// See (*database/sql.Rows).Close() for additional documentation.
func (*Rows) Close() {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L180-L180
	doPanic("encore apps must be run using the encore command")
	return
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
//
// See (*database/sql.Rows).Scan() for additional documentation.
func (*Rows) Scan(dest ...interface{}) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L187-L187
	doPanic("encore apps must be run using the encore command")
	return
}

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit Close.
//
// See (*database/sql.Rows).Err() for additional documentation.
func (*Rows) Err() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L193-L193
	doPanic("encore apps must be run using the encore command")
	return
}

// Next prepares the next result row for reading with the Scan method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. Err should be consulted to distinguish between
// the two cases.
//
// Every call to Scan, even the first one, must be preceded by a call to Next.
//
// See (*database/sql.Rows).Next() for additional documentation.
func (*Rows) Next() (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L203-L203
	doPanic("encore apps must be run using the encore command")
	return
}

// Row is the result of calling QueryRow to select a single row.
//
// See *database/sql.Row for additional documentation.
type Row struct {
	_ int // for godoc to show unexported fields
}

// Scan copies the columns from the matched row into the values
// pointed at by dest.
//
// See (*database/sql.Row).Scan() for additional documentation.
func (*Row) Scan(dest ...interface{}) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L217-L230
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Row) Err() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/storage/sqldb/sqldb.go#L232-L237
	doPanic("encore apps must be run using the encore command")
	return
}
