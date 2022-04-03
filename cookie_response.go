package gsr7

import (
	"fmt"
	"net/url"
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
	// WithDomain sets the domain name on the current cookie. The function returns an error if the domain is not valid
	// for cookies. See https://datatracker.ietf.org/doc/html/rfc6265#section-5.2.3 for details.
	WithDomain(domain string) (ResponseCookie, error)

	// GetPath returns the path
	GetPath() string
	WithPath(path string) (ResponseCookie, error)

	GetExpires() *time.Time
	WithExpires(expires *time.Time) ResponseCookie

	GetMaxAge() *int
	WithMaxAge(deltaSeconds *int) ResponseCookie

	GetSecure() bool
	WithSecure(secure bool) ResponseCookie

	GetHTTPOnly() bool
	WithHTTPOnly(httpOnly bool) ResponseCookie
}

func NewResponseCookie(name string) (ResponseCookie, error) {
	if err := validate(validateCookieName(name)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name: name,
	}, nil
}

//region Implementation
type responseCookie struct {
	name     string
	value    string
	path     string
	domain   string
	expires  *time.Time
	maxAge   *int
	secure   bool
	httpOnly bool
}

func (r responseCookie) Name() string {
	return r.name
}

func (r responseCookie) Value() string {
	return r.value
}

func (r responseCookie) WithName(name string) (ResponseCookie, error) {
	if err := validate(validateCookieName(name)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name:     name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   r.maxAge,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}, nil
}

func (r responseCookie) WithValue(value string) (ResponseCookie, error) {
	return &responseCookie{
		name:     r.name,
		value:    value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   r.maxAge,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}, nil
}

func (r responseCookie) GetDomain() string {
	return r.domain
}

func (r responseCookie) WithDomain(domain string) (ResponseCookie, error) {
	if err := validate(validateCookieDomain(domain)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   domain,
		expires:  r.expires,
		maxAge:   r.maxAge,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}, nil
}

func (r responseCookie) GetPath() string {
	return r.path
}

func (r responseCookie) WithPath(path string) (ResponseCookie, error) {
	if err := validate(validateCookiePath(path)); err != nil {
		return nil, err
	}
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   r.maxAge,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}, nil
}

func (r responseCookie) GetExpires() *time.Time {
	return r.expires
}

func (r responseCookie) WithExpires(expires *time.Time) ResponseCookie {
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  expires,
		maxAge:   r.maxAge,
		secure:   r.secure,
		httpOnly: r.httpOnly,
	}
}

func (r responseCookie) GetMaxAge() *int {
	return r.maxAge
}

func (r responseCookie) WithMaxAge(deltaSeconds *int) ResponseCookie {
	return &responseCookie{
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   deltaSeconds,
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
		name:     r.name,
		value:    r.value,
		path:     r.path,
		domain:   r.domain,
		expires:  r.expires,
		maxAge:   r.maxAge,
		secure:   r.secure,
		httpOnly: httpOnly,
	}
}

func (r responseCookie) Encode() string {
	result := fmt.Sprintf("%s=%s", r.name, url.QueryEscape(r.value))
	if r.path != "" {
		result = fmt.Sprintf("%s; path=%s", result, r.path)
	}
	if r.domain != "" {
		result = fmt.Sprintf("%s; domain=%s", result, r.domain)
	}
	if r.expires != nil {
		result = fmt.Sprintf("%s; expires=%s", result, r.expires.UTC().Format(time.RFC1123))
	}
	if r.secure {
		result = fmt.Sprintf("%s; secure", result)
	}
	if r.httpOnly {
		result = fmt.Sprintf("%s; httpOnly", result)
	}
	return result
}

//endregion
