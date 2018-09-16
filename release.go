// +build !debug

package errors

// Serializable returns a serializable error struct that cna omit exposing internal errors as an option in release mode
func Serializable(err error) *Error {
	e := Unwrap(err)
	e.cause = nil
	return e
}
