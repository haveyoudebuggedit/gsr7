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
	Name() string
	Value() string

	WithName(name string) T
	WithNameE(name string) (T, error)
	WithValue(value string) T
	WithValueE(value string) (T, error)

	// Encode encodes the current cookie into a format compatible with the Cookie or SetCookie header.
	Encode() string
}

//endregion
