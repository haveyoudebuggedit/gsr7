package gsr7

type Comparable[T any] interface {
	// Compare compares the current object to the other object and returns a negative number if the current object is
	// smaller than the other, a positive number if it is larger, and 0 if it is equal.
	Compare(other T) int
}
