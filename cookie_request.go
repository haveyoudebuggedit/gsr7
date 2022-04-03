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
}

// NewRequestCookie creates a new request cookie from the specified name and value
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

func (r requestCookie) Name() string {
	return r.name
}

func (r requestCookie) Value() string {
	return r.value
}

func (r requestCookie) WithName(name string) (RequestCookie, error) {
	if err := validate(validateCookieName(name)); err != nil {
		return nil, err
	}
	return &requestCookie{
		name:  name,
		value: r.value,
	}, nil
}

func (r requestCookie) WithValue(value string) (RequestCookie, error) {
	return &requestCookie{
		name:  r.name,
		value: value,
	}, nil
}

func (r requestCookie) Encode() string {
	return fmt.Sprintf("%s=%s", r.name, url.QueryEscape(r.value))
}

//endregion
