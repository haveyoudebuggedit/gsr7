package gsr7_test

import (
	"fmt"
	"testing"

	"go.debugged.it/gsr7"
)

//region Examples

func ExampleParseVersion() {
	version, err := gsr7.ParseVersion("HTTP/1.1")
	if err != nil {
		panic(err)
	}
	fmt.Println(version)

	// Output: HTTP/1.1
}

func ExampleNewVersion() {
	version := gsr7.Must(gsr7.NewVersion(0, 9))
	fmt.Println(version)

	// Output: HTTP/0.9
}

func ExampleVersion_Equals() {
	if gsr7.HTTP10.Equals(gsr7.HTTP11) {
		fmt.Println("HTTP/1.0 equals HTTP/1.1")
	} else {
		fmt.Println("HTTP/1.0 does not equal HTTP/1.1")
	}

	// Output: HTTP/1.0 does not equal HTTP/1.1
}

func ExampleVersion_Compare() {
	if gsr7.HTTP10.Compare(gsr7.HTTP11) < 0 {
		fmt.Println("HTTP/1.0 is lower than HTTP/1.1")
	} else {
		fmt.Println("HTTP/1.0 is not lower than HTTP/1.1")
	}

	// Output: HTTP/1.0 is lower than HTTP/1.1
}

func ExampleVersion_String() {
	fmt.Println(gsr7.HTTP11.String())

	// Output: HTTP/1.1
}

func ExampleVersion_Major() {
	fmt.Println(gsr7.HTTP10.Major())

	// Output: 1
}

func ExampleVersion_Minor() {
	fmt.Println(gsr7.HTTP10.Minor())

	// Output: 0
}

//endregion

//region Tests

func TestParseVersion(t *testing.T) {
	validData := map[string]gsr7.Version{
		"HTTP/0.9": gsr7.HTTP09,
		"HTTP/1.0": gsr7.HTTP10,
		"HTTP/1.1": gsr7.HTTP11,
		"HTTP/2.0": gsr7.HTTP20,
	}
	for input, expectedOutput := range validData {
		t.Run(
			input, func(t *testing.T) {
				ver, err := gsr7.ParseVersion(input)
				if err != nil {
					t.Fatalf("failed to parse %s (%v)", input, err)
				}
				if !ver.Equals(expectedOutput) {
					t.Fatalf(
						"resulting version of %s does not equal the expected version of %s",
						ver,
						expectedOutput,
					)
				}
			},
		)
	}

	invalidData := []string{
		"",
		"HTTP/",
		"HTTP/.1",
		"HTTP/1",
		"HTTP/1.",
		"HTTP/0.0",
	}
	for _, input := range invalidData {
		t.Run(
			input,
			func(t *testing.T) {
				_, err := gsr7.ParseVersion(input)
				if err == nil {
					t.Fatalf("parsing the string '%s' did not result in an error", input)
				}
			},
		)
	}
}

func TestVersionCompare(t *testing.T) {
	assertSmallerThan(
		t,
		gsr7.HTTP09.Compare(gsr7.HTTP10),
		0,
		"HTTP/0.9 is not smaller than HTTP/1.1",
	)
	assertSmallerThan(
		t,
		gsr7.HTTP10.Compare(gsr7.HTTP11),
		0,
		"HTTP/1.0 is not smaller than HTTP/1.1",
	)
	assertEquals(
		t,
		gsr7.HTTP10.Compare(gsr7.HTTP10),
		0,
		"HTTP/1.0 is not equal to HTTP/1.0",
	)
	assertEquals(
		t,
		gsr7.HTTP11.Compare(gsr7.HTTP11),
		0,
		"HTTP/1.0 is not equal to HTTP/1.0",
	)
}

//endregion
