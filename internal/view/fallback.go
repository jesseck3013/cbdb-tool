package view

import "fmt"

const FALLBACK = "Unkown"

func valueOrFallBack[T any](ptr *T, fallback string) string {
	if ptr == nil {
		return fallback
	}

	val, ok := any(*ptr).(string)
	if ok && val == "" {
		return fallback
	}

	return fmt.Sprintf("%v", *ptr)
}

func safeValue[T any](ptr *T) string {
	return valueOrFallBack(ptr, FALLBACK)
}
