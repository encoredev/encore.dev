package sqldb

import (
	"encore.dev/storage/sqldb/sqlerr"
)

// ErrCode reports the error code for a given error.
// If the error is nil or is not of type *Error it reports sqlerr.Other.
func ErrCode(err error) (_ sqlerr.Code) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/sqldb/errors.go#L16-L22
	doPanic("encore apps must be run using the encore command")
	return
}

// Error represents an error reported by the database server.
// It's not guaranteed all errors reported by sqldb functions will be of this type;
// it is only returned when the database reports an error.
//
// Note that the fields for schema name, table name, column name, data type name,
// and constraint name are supplied only for a limited number of error types;
// see https://www.postgresql.org/docs/current/errcodes-appendix.html.
//
// You should not assume that the presence of any of these fields guarantees
// the presence of another field.
type Error struct {
	// Code defines the general class of the error.
	// See [sqlerr.Code] for a list of possible values.
	Code sqlerr.Code

	// Severity is the severity of the error.
	Severity sqlerr.Severity

	// DatabaseCode is the database server-specific error code.
	// It is specific to the underlying database server.
	DatabaseCode string

	// Message: the primary human-readable error message. This should be accurate
	// but terse (typically one line). Always present.
	Message string

	// SchemaName: if the error was associated with a specific database object,
	// the name of the schema containing that object, if any.
	SchemaName string

	// TableName: if the error was associated with a specific table, the name of the table.
	// (Refer to the schema name field for the name of the table's schema.)
	TableName string

	// ColumnName: if the error was associated with a specific table column,
	// the name of the column. (Refer to the schema and table name fields to identify the table.)
	ColumnName string

	// Data type name: if the error was associated with a specific data type,
	// the name of the data type. (Refer to the schema name field for the name of the data type's schema.)
	DataTypeName string

	// Constraint name: if the error was associated with a specific constraint,
	// the name of the constraint. Refer to fields listed above for the associated
	// table or domain. (For this purpose, indexes are treated as constraints,
	// even if they weren't created with constraint syntax.)
	ConstraintName string
}

func (*Error) Error() (_ string) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/sqldb/errors.go#L78-L80
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Error) Unwrap() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/sqldb/errors.go#L82-L84
	doPanic("encore apps must be run using the encore command")
	return
}
