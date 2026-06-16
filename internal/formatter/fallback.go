package formatter

import "fmt"

func valueOrFallBack[T any](ptr *T, fallback string) string {
	if ptr == nil {
		return fallback
	}

	val, ok := any(*ptr).(string)
	if ok && val == "" {
		return val
	}

	return fmt.Sprintf("%v", *ptr)
}
