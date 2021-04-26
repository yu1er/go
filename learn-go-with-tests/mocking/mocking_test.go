package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("spy sleeper for count", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		// sleeper := &ConfigurableSleeper{time.Second}
		CounterDown(buffer, sleeper)
		got := buffer.String()
		expected := `3
2
1
Go!`

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}

		if sleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
		}
	})

	t.Run("configurable sleeper for time sleep", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &ConfigurableSleeper{time.Second}
		CounterDown(buffer, sleeper)
		got := buffer.String()
		expected := `3
2
1
Go!`

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("counter down sleeper for the order of sleep and write", func(t *testing.T) {
		sleeper := &CounterDownSleeper{}
		CounterDown(sleeper, sleeper)
		got := sleeper.Calls
		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

}
