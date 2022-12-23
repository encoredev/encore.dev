package sqldb

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	_ int // for godoc to show unexported fields
}

// Stdlib returns a *sql.DB object that is connected to the same db,
// for use with libraries that expect a *sql.DB.
func (*Database) Stdlib() *sql.DB {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/storage/sqldb/db.go#L46-L77
	panic("encore apps must be run using the encore command")
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).ExecContext() for additional documentation.
func (*Database) Exec(ctx context.Context, query string, args ...interface{}) (ExecResult, error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/storage/sqldb/db.go#L151-L175
	panic("encore apps must be run using the encore command")
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).QueryContext() for additional documentation.
func (*Database) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/storage/sqldb/db.go#L181-L208
	panic("encore apps must be run using the encore command")
}

// QueryRow executes a query that is expected to return at most one row.
//
// See (*database/sql.DB).QueryRowContext() for additional documentation.
func (*Database) QueryRow(ctx context.Context, query string, args ...interface{}) *Row {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/storage/sqldb/db.go#L213-L238
	panic("encore apps must be run using the encore command")
}

// Begin opens a new database transaction.
//
// See (*database/sql.DB).Begin() for additional documentation.
func (*Database) Begin(ctx context.Context) (*Tx, error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/storage/sqldb/db.go#L243-L263
	panic("encore apps must be run using the encore command")
}

// Driver returns the underlying database driver for this database connection pool.
//
//	var db = sqldb.Driver[*pgxpool.Pool](sqldb.Named("mydatabase"))
//
// This is defined as a generic function to allow compile-time type checking
// that the Encore application is expecting a driver that is supported.
//
// At some point in the future where Encore adds support for a different database driver
// this will be made with backwards compatibility in mind, providing ample notice and
// time to migrate in an opt-in fashion.
func Driver[T SupportedDrivers](db *Database) T {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/storage/sqldb/db.go#L275-L277
	panic("encore apps must be run using the encore command")
}

// SupportedDrivers is a type list of all supported database drivers.
// Currently only [*pgxpool.Pool] is supported.
type SupportedDrivers interface {
	*pgxpool.Pool
}
