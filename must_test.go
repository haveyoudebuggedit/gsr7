package gsr7_test

import (
	"fmt"

	"go.debugged.it/gsr7"
)

// ExampleMust showcases how the Must function can be used to convert errors to panics. This is useful in cases where
// you know that no error will occur.
func ExampleMust() {
	ver := gsr7.Must(gsr7.ParseVersion("HTTP/1.1"))
	fmt.Println(ver.String())

	// Output: HTTP/1.1
}
