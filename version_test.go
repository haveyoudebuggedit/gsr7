package gsr7_test

import (
	"fmt"
	"sort"
	"testing"

	"go.debugged.it/gsr7"
)

func ExampleParseVersion() {
	version, err := gsr7.ParseVersion("HTTP/1.1")
	if err != nil {
		panic(err)
	}
	fmt.Println(version)

	// Output: HTTP/1.1
}

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
	list := []gsr7.Version{
		gsr7.HTTP20,
		gsr7.HTTP11,
		gsr7.HTTP10,
		gsr7.HTTP09,
	}
	sort.Slice(
		list, func(i, j int) bool {
			return list[i].Compare(list[j]) > 0
		},
	)

	if !list[0].Equals(gsr7.HTTP09) {
		t.Fatalf("incorrect list sorting (item 0 is not HTTP/0.9)")
	}
	if !list[1].Equals(gsr7.HTTP10) {
		t.Fatalf("incorrect list sorting (item 0 is not HTTP/1.0)")
	}
	if !list[2].Equals(gsr7.HTTP11) {
		t.Fatalf("incorrect list sorting (item 0 is not HTTP/1.1)")
	}
	if !list[3].Equals(gsr7.HTTP20) {
		t.Fatalf("incorrect list sorting (item 0 is not HTTP/2.0)")
	}
}
