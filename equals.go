package gsr7

// Equals defines the Equals method to compare the current object to another object.
type Equals[T any] interface {
	Equals(other T) bool
}
