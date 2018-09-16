// +build debug

package errors

// Return a serializable error struct that cna omit exposing internal errors as an option in release mode
func Serializable(err error) *Error {
	return Unwrap(err)
}
