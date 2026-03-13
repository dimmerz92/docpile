package core

// Coalesce returns the first non-zero value in values for type T.
// If no non-zero value is found, the zero value for type T is returned.
func Coalesce[T comparable](values ...T) T {
	var zero T

	for _, value := range values {
		if value != zero {
			return value
		}
	}

	return zero
}

// IIF or inline-if that returns v1 if the condition is true, otherwise false.
func IIF[T any](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}

	return v2
}
