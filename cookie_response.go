package gsr7

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// ResponseCookie is a cookie with additional information on cookie handling. This is sent to the client using the
// Set-Cookie header.
//
// See https://datatracker.ietf.org/doc/html/rfc6265#section-4.1 for details.
type ResponseCookie interface {
	Cookie[ResponseCookie]

	// GetDomain returns the domain name (if set). This doesn't have to be a valid FQDN as it can start with a dot
	// indicating all subdomains. See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.3 for details.
	GetDomain() string
	// WithDomain returns a new cookie object valid for the specified domain. This function does not modify the
	// original cookie. The function panics if the domain is not valid for cookies.
	//
	// See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.3 for details.
	WithDomain(domain string) ResponseCookie
	// WithDomainE returns a new cookie object valid for the specified domain. This function does not modify the
	// original cookie. The function returns an error if the domain is not valid for cookies.
	//
	// See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.3 for details.
	WithDomainE(domain string) (ResponseCookie, error)

	// GetPath returns the path the cookie is valid for.
	// See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.4 for details.
	GetPath() string
	// WithPath returns a new cookie object valid for the specified path. This function does not modify the original
	// cookie. The function panics if the path is not valid for cookies.
	//
	// See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.4 for details.
	WithPath(path string) ResponseCookie
	// WithPathE returns a new cookie object valid for the specified path. This function does not modify the original
	// cookie. The function returns an error if the path is not valid for cookies.
	//
	// See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.4 for details.
	WithPathE(path string) (ResponseCookie, error)

	// GetExpires returns the expiry time for the cookie, if set. If the expiry time is in the past the client will
	// delete the cookie.
	GetExpires() *time.Time
	// WithExpires returns a copy of the cookie with the expiry time set. If the expiry time is in the past the client
	// will delete the cookie.
	WithExpires(expires time.Time) ResponseCookie
	// WithoutExpires returns a copy of the cookie with the expiry time removed.
	WithoutExpires() ResponseCookie

	// GetMaxAge returns either nil, if max-age is not set, or the max-age for the cookie. A negative max-age will cause
	// the client to delete the cookie.
	GetMaxAge() *int
	// WithMaxAge returns a new cookie with the set max-age parameter. A negative max-age will cause the client to
	// delete the cookie.
	WithMaxAge(deltaSeconds int) ResponseCookie
	// WithoutMaxAge returns a cookie with the max-age parameter removed.
	WithoutMaxAge() ResponseCookie

	// GetSecure returns true if the cookie should only be transmitted over a secure (HTTPS) connection.
	GetSecure() bool
	// WithSecure creates a modified cookie with the secure flag set or unset. This flag indicates to the browser that
	// the cookie should only be transmitted over a secure (HTTPS) connection.
	WithSecure(secure bool) ResponseCookie

	// GetHTTPOnly returns true if the cookie should not be exposed to JavaScript code.
	GetHTTPOnly() bool
	// WithHTTPOnly returns a modified cookie with the httpOnly flag set or unset. This flag indicates to the browser
	// that the cookie should not be exposed to JavaScript.
	WithHTTPOnly(httpOnly bool) ResponseCookie

	// GetExtensions returns the extensions appended to the Set-Cookie header that are not otherwise known by
	// this implementation. Extensions may contain any character except for control characters (ASCII 0-27, 127) and
	// semicolons.
	GetExtensions() []string
	// WithExtensions creates a new cookie with the specified extensions. Extensions are other tags attached with a
	// semicolon that are not part of the known standards for cookies. Extensions may contain any character except for
	// control characters (ASCII 0-27, 127) and semicolons. If the extensions are invalid a panic is thrown.
	WithExtensions(extensions []string) ResponseCookie
	// WithExtensionsE creates a new cookie with the specified extensions. Extensions are other tags attached with a
	// semicolon that are not part of the known standards for cookies. Extensions may contain any character except for
	// control characters (ASCII 0-27, 127) and semicolons. If the extensions are invalid an error is returned.
	WithExtensionsE(extensions []string) (ResponseCookie, error)

	ToRequest() RequestCookie
}

// NewResponseCookie creates a new response cookie with the specified name. Other parameters can be added using
//// methods on the returned cookie. If the name is invalid a panic is thrown.
func NewResponseCookie(name string) ResponseCookie {
	return Must(NewResponseCookieE(name))
}

// NewResponseCookieE creates a new response cookie with the specified name. Other parameters can be added using
// methods on the returned cookie. If the name is invalid an error is returned.
func NewResponseCookieE(name string) (ResponseCookie, error) {
	if err := validate(validateCookieName(name)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name: name,
	}, nil
}

//region Implementation
type responseCookie struct {
	name       string
	value      string
	path       string
	domain     string
	expires    *time.Time
	maxAge     *int
	secure     bool
	httpOnly   bool
	extensions []string
}

func (r responseCookie) GetExtensions() []string {
	target := make([]string, len(r.extensions))
	copy(target, r.extensions)
	return target
}

func (r responseCookie) WithExtensions(extensions []string) ResponseCookie {
	return Must(r.WithExtensionsE(extensions))
}

func (r responseCookie) WithExtensionsE(extensions []string) (ResponseCookie, error) {
	if err := validate(validateExtensions(extensions)); err != nil {
		return nil, err
	}
	newExtensions := make([]string, len(extensions))
	copy(newExtensions, extensions)
	return &responseCookie{
		name:       r.name,
		value:      r.value,
		path:       r.path,
		domain:     r.domain,
		expires:    r.expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: newExtensions,
	}, nil
}

func (r responseCookie) ToRequest() RequestCookie {
	return &requestCookie{
		name:  r.name,
		value: r.value,
	}
}

func (r responseCookie) Name() string {
	return r.name
}

func (r responseCookie) Value() string {
	return r.value
}

func (r responseCookie) WithName(name string) ResponseCookie {
	return Must(r.WithNameE(name))
}

func (r responseCookie) WithNameE(name string) (ResponseCookie, error) {
	if err := validate(validateCookieName(name)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name:       name,
		value:      r.value,
		path:       r.path,
		domain:     r.domain,
		expires:    r.expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: r.extensions,
	}, nil
}

func (r responseCookie) WithValue(value string) ResponseCookie {
	return Must(r.WithValueE(value))
}

func (r responseCookie) WithValueE(value string) (ResponseCookie, error) {
	return &responseCookie{
		name:       r.name,
		value:      value,
		path:       r.path,
		domain:     r.domain,
		expires:    r.expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: r.extensions,
	}, nil
}

func (r responseCookie) GetDomain() string {
	return r.domain
}

func (r responseCookie) WithDomain(domain string) ResponseCookie {
	return Must(r.WithDomainE(domain))
}

func (r responseCookie) WithDomainE(domain string) (ResponseCookie, error) {
	if err := validate(validateCookieDomain(domain)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name:       r.name,
		value:      r.value,
		path:       r.path,
		domain:     domain,
		expires:    r.expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: r.extensions,
	}, nil
}

func (r responseCookie) GetPath() string {
	return r.path
}

func (r responseCookie) WithPath(path string) ResponseCookie {
	return Must(r.WithPathE(path))
}

func (r responseCookie) WithPathE(path string) (ResponseCookie, error) {
	if err := validate(validateCookiePath(path)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name:       r.name,
		value:      r.value,
		path:       path,
		domain:     r.domain,
		expires:    r.expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: r.extensions,
	}, nil
}

func (r responseCookie) GetExpires() *time.Time {
	return r.expires
}

func (r responseCookie) WithExpires(expires time.Time) ResponseCookie {
	return &responseCookie{
		name:       r.name,
		value:      r.value,
		path:       r.path,
		domain:     r.domain,
		expires:    &expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: r.extensions,
	}
}

func (r responseCookie) WithoutExpires() ResponseCookie {
	return &responseCookie{
		name:       r.name,
		value:      r.value,
		path:       r.path,
		domain:     r.domain,
		expires:    nil,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   r.httpOnly,
		extensions: r.extensions,
	}
}

func (r responseCookie) GetMaxAge() *int {
	return r.maxAge
}

func (r responseCookie) WithMaxAge(deltaSeconds int) ResponseCookie {
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   &deltaSeconds,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}
}

func (r responseCookie) WithoutMaxAge() ResponseCookie {
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   nil,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}
}

func (r responseCookie) GetSecure() bool {
	return r.secure
}

func (r responseCookie) WithSecure(secure bool) ResponseCookie {
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   r.maxAge,
		secure:   secure,
		httpOnly: r.httpOnly,
	}
}

func (r responseCookie) GetHTTPOnly() bool {
	return r.httpOnly
}

func (r responseCookie) WithHTTPOnly(httpOnly bool) ResponseCookie {
	return &responseCookie{
		name:       r.name,
		value:      r.value,
		path:       r.path,
		domain:     r.domain,
		expires:    r.expires,
		maxAge:     r.maxAge,
		secure:     r.secure,
		httpOnly:   httpOnly,
		extensions: r.extensions,
	}
}

func (r responseCookie) Encode() string {
	parts := []string{
		fmt.Sprintf("%s=%s", r.name, url.QueryEscape(r.value)),
	}
	if r.path != "" {
		parts = append(parts, fmt.Sprintf("path=%s", r.path))
	}
	if r.domain != "" {
		parts = append(parts, fmt.Sprintf("domain=%s", r.domain))
	}
	if r.expires != nil {
		parts = append(parts, fmt.Sprintf("expires=%s", r.expires.UTC().Format(time.RFC1123)))
	}
	if r.secure {
		parts = append(parts, "secure")
	}
	if r.httpOnly {
		parts = append(parts, "httpOnly")
	}
	parts = append(parts, r.extensions...)
	return strings.Join(parts, "; ")
}

//endregion
