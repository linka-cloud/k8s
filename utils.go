package k8s

// Ref returns a pointer to the given value.
func Ref[T any](t T) *T {
	return &t
}

// Value returns the value of the given pointer, or the default value if the pointer is nil.
func Value[T any](t *T, def ...T) (v T) {
	if t != nil {
		return *t
	}
	if len(def) > 0 {
		return def[0]
	}
	return
}

func Equal[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
