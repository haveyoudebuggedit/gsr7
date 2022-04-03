package gsr7

// Must provides error-to-exception conversion on return values.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
