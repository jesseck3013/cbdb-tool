package formatter

import "testing"

func assertString(t *testing.T, got string, expect string) {
	t.Helper()
	if got != expect {
		t.Errorf("got %s want %s", got, expect)
	}
}

func TestFallBack(t *testing.T) {
	unknown := "Unkown"
	t.Run("non nil *int16", func(t *testing.T) {
		intValue := 10
		got := valueOrFallBack(&intValue, unknown)
		assertString(t, got, "10")
	})

	t.Run("nil *int16", func(t *testing.T) {
		var nilInt *int16
		got := valueOrFallBack(nilInt, unknown)
		assertString(t, got, unknown)
	})

	t.Run("non nil *string", func(t *testing.T) {
		stringValue := "test"
		got := valueOrFallBack(&stringValue, unknown)
		assertString(t, got, stringValue)
	})

	t.Run("non nil empty string", func(t *testing.T) {
		stringValue := ""
		got := valueOrFallBack(&stringValue, unknown)
		assertString(t, got, unknown)
	})

	t.Run("nil *string", func(t *testing.T) {
		var stringPtr *string
		got := valueOrFallBack(stringPtr, unknown)
		assertString(t, got, unknown)
	})
}

func TestSafeValue(t *testing.T) {
	t.Run("non nil *int16", func(t *testing.T) {
		intValue := 10
		got := safeValue(&intValue)
		assertString(t, got, "10")
	})

	t.Run("nil *int16", func(t *testing.T) {
		var nilInt *int16
		got := safeValue(nilInt)
		assertString(t, got, FALLBACK)
	})

	t.Run("non nil *string", func(t *testing.T) {
		stringValue := "test"
		got := safeValue(&stringValue)
		assertString(t, got, stringValue)
	})

	t.Run("non nil empty string", func(t *testing.T) {
		stringValue := ""
		got := safeValue(&stringValue)
		assertString(t, got, FALLBACK)
	})

	t.Run("nil *string", func(t *testing.T) {
		var stringPtr *string
		got := safeValue(stringPtr)
		assertString(t, got, FALLBACK)
	})
}
