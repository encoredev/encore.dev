package sqldb

import (
	"context"
	"database/sql"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	_ int // for godoc to show unexported fields
}

// Hooks defines callbacks that can be registered for database lifecycle events.
type Hooks struct {
	// AfterConnect is called whenever a new database connection is established.
	// Returning an error aborts creation of that connection.
	AfterConnect func(context.Context, *pgx.Conn) error
}

// AddHooks registers callbacks to run for this database.
// Does not have any effect when using the database/sql integration via the (*Database).Stdlib method.
func (*Database) AddHooks(h Hooks) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L55-L59
	doPanic("encore apps must be run using the encore command")
	return
}

// Stdlib returns a *sql.DB object that is connected to the same db,
// for use with libraries that expect a *sql.DB.
func (*Database) Stdlib() (_ *sql.DB) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L101-L134
	doPanic("encore apps must be run using the encore command")
	return
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).ExecContext() for additional documentation.
func (*Database) Exec(ctx context.Context, query string, args ...interface{}) (_ ExecResult, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L207-L243
	doPanic("encore apps must be run using the encore command")
	return
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
//
// See (*database/sql.DB).QueryContext() for additional documentation.
func (*Database) Query(ctx context.Context, query string, args ...interface{}) (_ *Rows, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L249-L287
	doPanic("encore apps must be run using the encore command")
	return
}

// QueryRow executes a query that is expected to return at most one row.
//
// See (*database/sql.DB).QueryRowContext() for additional documentation.
func (*Database) QueryRow(ctx context.Context, query string, args ...interface{}) (_ *Row) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L292-L328
	doPanic("encore apps must be run using the encore command")
	return
}

// Begin opens a new database transaction.
//
// See (*database/sql.DB).Begin() for additional documentation.
func (*Database) Begin(ctx context.Context) (_ *Tx, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L333-L356
	doPanic("encore apps must be run using the encore command")
	return
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
func Driver[T SupportedDrivers](db *Database) (_ T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L368-L376
	doPanic("encore apps must be run using the encore command")
	return
}

// SupportedDrivers is a type list of all supported database drivers.
// Currently only [*pgxpool.Pool] is supported.
type SupportedDrivers interface {
	*pgxpool.Pool
}

// DriverConn provides access to the underlying driver connection given a stdlib
// *sql.Conn connection. The driverConn must not be used outside of f, and conn
// must be a *sql.Conn originating from sqldb.Database or this function panics.
//
//	conn, _ := db.Stdlib().Conn(ctx) // Checkout a connection from the pool
//	sqldb.DriverConn(conn, func(driverConn *pgx.Conn) error) error {
//	  // do stuff with *pgx.Conn
//	}
//
// This is defined as a generic function to allow compile-time type checking
// that the Encore application is expecting a driver that is supported.
//
// At some point in the future where Encore adds support for a different
// database driver this will be made with backwards compatibility in mind,
// providing ample notice and time to migrate in an opt-in fashion.
func DriverConn[T SupportedDriverConns](conn *sql.Conn, f func(driverConn T) error) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/sqldb/db.go#L399-L412
	doPanic("encore apps must be run using the encore command")
	return
}

// SupportedDriverConns is a type list of all supported database drivers
// connections. Currently only [*pgx.Conn] is supported.
type SupportedDriverConns interface {
	*pgx.Conn
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking Encore APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("ENCORERUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
