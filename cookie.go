package gsr7

import "time"

type cookie[T cookie] interface {
    Name() string
    Value() string

    WithName(name string) (T, error)
    WithValue(value string) (T, error)
}

type RequestCookie interface {
    cookie[RequestCookie]
}

type ResponseCookie interface {
    cookie[ResponseCookie]

    GetDomain() string
    WithDomain(domain string) (ResponseCookie, error)

    GetPath() string
    WithPath(path string) (ResponseCookie, error)

    GetExpires() time.Time
    WithExpires(expires time.Time) ResponseCookie

    GetMaxAge() *int
    SetMaxAge(deltaSeconds *int) ResponseCookie

    GetSecure() bool
    WithSecure(secure bool) ResponseCookie

    GetHTTPOnly() bool
    WithHTTPOnly(httpOnly bool) ResponseCookie
}
