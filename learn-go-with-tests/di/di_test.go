package di

import (
	"bytes"
	"testing"
)

func TestDi(t *testing.T) {
	t.Run("di", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Greet(buffer, "Chris")

		got := buffer.String()
		expected := "Hello, Chris"

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
