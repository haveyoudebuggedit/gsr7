package gsr7_test

import (
	"fmt"
	"time"

	"go.debugged.it/gsr7"
)

func ExampleResponseCookie() {
	cookie := gsr7.
		NewResponseCookie("foo").
		WithValue("bar").
		WithDomain("example.com").
		WithPath("/").
		WithExpires(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Printf("Set-Cookie: %s", cookie.Encode())
	// Output: Set-Cookie: foo=bar; path=/; domain=example.com; expires=Thu, 01 Jan 1970 00:00:00 UTC
}
