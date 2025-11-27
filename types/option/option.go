// Package option provides a generic Option type for representing optional values
// in a more type-safe way than using pointers or zero values.
package option

import "os"

// Option is a type that represents a value that may or may not be present.
type Option[T any] struct {
	_ int // for godoc to show unexported fields
}

func (Option[T]) MarshalJSON() (_ []byte, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L16-L21
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Option[T]) UnmarshalJSON(data []byte) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L23-L33
	doPanic("encore apps must be run using the encore command")
	return
}

// FromComparable returns Some(v) if v is not zero, and None otherwise.
// If T implements an IsZero() bool method, that is also used to determine if v is zero.
func FromComparable[T comparable](v T) (_ Option[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L37-L45
	doPanic("encore apps must be run using the encore command")
	return
}

// FromPointer returns Some(*v) if v is not nil, and None otherwise.
func FromPointer[T any](v *T) (_ Option[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L48-L53
	doPanic("encore apps must be run using the encore command")
	return
}

// Some returns an Option with the given value and present set to true.
func Some[T any](v T) (_ Option[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L56-L58
	doPanic("encore apps must be run using the encore command")
	return
}

// None returns an Option with no value set.
func None[T any]() (_ Option[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L61-L63
	doPanic("encore apps must be run using the encore command")
	return
}

// IsSome returns true if the Option has a value set.
func (Option[T]) IsSome() (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L66-L68
	doPanic("encore apps must be run using the encore command")
	return
}

// IsNone returns true if the Option has no value set.
func (Option[T]) IsNone() (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L71-L73
	doPanic("encore apps must be run using the encore command")
	return
}

// IsZero is an alias for IsNone, to support usage in structs with "omitempty".
func (Option[T]) IsZero() (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L76-L78
	doPanic("encore apps must be run using the encore command")
	return
}

// Get gets the option value and returns ok==true if present.
// Commonly used in the "comma ok" idiom:
//
//	if val, ok := option.Get(); ok {
//	    ...
//	}
func (Option[T]) Get() (val T, ok bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L86-L88
	doPanic("encore apps must be run using the encore command")
	return
}

// GetOrElse returns the value if present, otherwise returns alternative.
func (Option[T]) GetOrElse(alternative T) (_ T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L91-L96
	doPanic("encore apps must be run using the encore command")
	return
}

// GetOrElseF returns the value if present, otherwise returns alternative().
func (Option[T]) GetOrElseF(alternative func() T) (_ T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L99-L104
	doPanic("encore apps must be run using the encore command")
	return
}

// MustGet returns the value if present, and otherwise panics.
func (Option[T]) MustGet() (_ T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L107-L112
	doPanic("encore apps must be run using the encore command")
	return
}

// OrElse returns an Option with the value if present, otherwise returns Some(alternative).
func (Option[T]) OrElse(alternative T) (_ Option[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L115-L120
	doPanic("encore apps must be run using the encore command")
	return
}

// Contains returns pred(v) if the option contains v, and false otherwise.
func (Option[T]) Contains(predicate func(v T) bool) (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L123-L128
	doPanic("encore apps must be run using the encore command")
	return
}

func (Option[T]) String() (_ string) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L130-L135
	doPanic("encore apps must be run using the encore command")
	return
}

func (Option[T]) GoString() (_ string) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L137-L142
	doPanic("encore apps must be run using the encore command")
	return
}

// PtrOrNil returns the value as a pointer if present, or nil otherwise.
func (Option[T]) PtrOrNil() (_ *T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L145-L150
	doPanic("encore apps must be run using the encore command")
	return
}

// Equal reports whether a and b are equal, using ==.
// If both are None, they are considered equal.
func Equal[T comparable](a, b Option[T]) (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L154-L162
	doPanic("encore apps must be run using the encore command")
	return
}

// Contains returns true if the option is present and matches the given value.
func Contains[T comparable](option Option[T], matches T) (_ bool) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L165-L170
	doPanic("encore apps must be run using the encore command")
	return
}

// Map returns an Option with the value mapped by the given function if present, otherwise returns None.
func Map[T, R any](option Option[T], f func(T) R) (_ Option[R]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/types/option/option.go#L173-L178
	doPanic("encore apps must be run using the encore command")
	return
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking Encore APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("ENCORERUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
