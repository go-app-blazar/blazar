package deref

import "fmt"

// String dereferences a pointer and returns a string representation of the value.
// If the value is nil, then this returns "nil".
func String[T any](value *T) string {
	if value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", *value)
}
