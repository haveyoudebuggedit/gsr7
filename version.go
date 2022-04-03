package gsr7

import (
	"fmt"
	"strconv"
	"strings"
)

// Version holds the HTTP
type Version interface {
	Equals[Version]
	Comparable[Version]

	// String must convert the version back to its HTTP header form of HTTP/major.minor
	String() string
	// Major returns the major version of the HTTP version.
	Major() uint8
	// Minor returns the minor version of the HTTP version.
	Minor() uint8
}

var HTTP09 = Must(NewVersion(0, 9))
var HTTP10 = Must(NewVersion(1, 0))
var HTTP11 = Must(NewVersion(1, 1))
var HTTP20 = Must(NewVersion(2, 0))

// NewVersion constructs a Version structure from the specified major and minor version.
func NewVersion(major, minor uint8) (Version, error) {
	if major == 0 && minor == 0 {
		return nil, fmt.Errorf("invalid HTTP version: %d.%d", major, minor)
	}
	return &version{
		major: major,
		minor: minor,
	}, nil
}

// ParseVersion parses a HTTP version string starting with HTTP/ into a version structure.
// Optionally, it can also ignore the HTTP/ prefix. If the version string is not valid it returns an error.
func ParseVersion(versionString string) (Version, error) {
	parts := strings.Split(versionString, "/")
	versionText := ""
	switch len(parts) {
	case 1:
		versionText = versionString
	case 2:
		if parts[0] != "HTTP" {
			return nil, fmt.Errorf("invalid HTTP version: %s", versionString)
		}
		versionText = parts[1]
	default:
		return nil, fmt.Errorf("invalid HTTP version: %s", versionString)
	}

	versionParts := strings.Split(versionText, ".")
	if len(versionParts) != 2 {
		return nil, fmt.Errorf("invalid HTTP version: %s", versionString)
	}
	major, err := strconv.Atoi(versionParts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP version: %s (%w)", versionParts, err)
	}
	if major > 255 {
		return nil, fmt.Errorf("invalid HTTP version: %s", versionParts)
	}
	minor, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP version: %s (%w)", versionParts, err)
	}
	if minor > 255 {
		return nil, fmt.Errorf("invalid HTTP version: %s", versionParts)
	}
	return NewVersion(uint8(major), uint8(minor))
}

// region Implementation

type version struct {
	major, minor uint8
}

func (v version) Compare(other Version) int {
	if v.major != other.Major() {
		return int(v.major) - int(other.Major())
	}
	return int(v.minor) - int(other.Minor())
}

func (v version) Equals(other Version) bool {
	return v.Compare(other) == 0
}

func (v version) String() string {
	return fmt.Sprintf("HTTP/%d.%d", v.major, v.minor)
}

func (v version) Major() uint8 {
	return v.major
}

func (v version) Minor() uint8 {
	return v.minor
}

// endregion
