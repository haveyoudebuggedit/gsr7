package gsr7

import (
	"fmt"
	"net/url"
)

// RequestCookie is a cookie that is set in a request using the Cookie header. It only contains the name and value
// of the cookie. A single cookie can be repeated multiple times and it is up to the server to figure out which one
// to use.
//
// See https://datatracker.ietf.org/doc/html/rfc6265#section-4.2.2 for details.
type RequestCookie interface {
	Cookie[RequestCookie]

	// ToResponse creates a response cookie from this request cookie. The additional parameters can be set on the
	// returned cookie.
	ToResponse() ResponseCookie
}

// NewRequestCookie creates a new request cookie from the specified name. The value can be added using the WithValue
// method.
func NewRequestCookie(name string) (RequestCookie, error) {
	if err := validate(
		validateCookieName(name),
	); err != nil {
		return nil, err
	}
	return &requestCookie{
		name,
		"",
	}, nil
}

//region Implementation

type requestCookie struct {
	name  string
	value string
}

func (r requestCookie) ToResponse() ResponseCookie {
	return &responseCookie{
		name:  r.name,
		value: r.value,
	}
}

func (r requestCookie) Name() string {
	return r.name
}

func (r requestCookie) Value() string {
	return r.value
}

func (r requestCookie) WithName(name string) RequestCookie {
	return Must(r.WithNameE(name))
}

func (r requestCookie) WithNameE(name string) (RequestCookie, error) {
	if err := validate(validateCookieName(name)); err != nil {
		return nil, err
	}
	return &requestCookie{
		name:  name,
		value: r.value,
	}, nil
}

func (r requestCookie) WithValue(value string) RequestCookie {
	return Must(r.WithValueE(value))
}

func (r requestCookie) WithValueE(value string) (RequestCookie, error) {
	return &requestCookie{
		name:  r.name,
		value: value,
	}, nil
}

func (r requestCookie) Encode() string {
	return fmt.Sprintf("%s=%s", r.name, url.QueryEscape(r.value))
}

//endregion
