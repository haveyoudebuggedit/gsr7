package gsr7

//region Interface

// Cookie is an interface definition for a single cookie as defined in RFC6265. This interface is the parent interface
// for both request and response cookies. As such, it only contains the name and value for the cookie.
//
// The actual implementations should use RequestCookie and ResponseCookie for more accurate typing. The T type
// constraint should be limited to Cookie, but Go doesn't support that.
//
// See https://datatracker.ietf.org/doc/html/rfc6265 for details.
type Cookie[T any] interface {
	// Name returns the name for the cookie.
	Name() string
	// Value returns the value for the cookie. This may be empty.
	Value() string

	// WithName creates a copy of the cookie with the specified name. If the name is invalid a panic is thrown.
	WithName(name string) T
	// WithNameE creates a copy of the cookie with the specified name. If the name is invalid an error is returned.
	WithNameE(name string) (T, error)
	// WithValue creates a copy of the cookie with the specified value. If the name is invalid a panic is thrown.
	WithValue(value string) T
	// WithValueE creates a copy of the cookie with the specified value. If the name is invalid an error is returned.
	WithValueE(value string) (T, error)

	// Encode encodes the current cookie into a format compatible with the Cookie or SetCookie header.
	Encode() string
}

//endregion
