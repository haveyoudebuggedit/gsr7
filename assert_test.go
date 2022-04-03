package gsr7_test

import (
	"testing"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func assertEquals[T comparable](t *testing.T, value, expected T, message string, args ...interface{}) {
	if value != expected {
		t.Fatalf(message, args...)
	}
}

func assertLargerThan[T Ordered](t *testing.T, value, expected T, message string, args ...interface{}) {
	if value <= expected {
		t.Fatalf(message, args...)
	}
}

func assertSmallerThan[T Ordered](t *testing.T, value, expected T, message string, args ...interface{}) {
	if value >= expected {
		t.Fatalf(message, args...)
	}
}
