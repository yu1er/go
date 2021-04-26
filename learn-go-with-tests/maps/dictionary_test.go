package maps

import "testing"

func TestSearch(t *testing.T) {

	dict := Dictionary{"test": "this is just a test"}

	t.Run("known", func(t *testing.T) {
		got, _ := dict.Search("test")
		expected := "this is just a test"
		assertStrings(t, expected, got)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := dict.Search("unknown key")

		assertError(t, ErrorNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{}
		dict.Add(key, value)

		assertValue(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}

		err := dict.Add(key, "another test")
		assertError(t, ErrorKeyExists, err)
		assertValue(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update test", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}
		newValue := "new test"
		dict.Update(key, newValue)
		assertValue(t, dict, key, newValue)
	})
}

func assertValue(t *testing.T, dict Dictionary, key, expected string) {
	t.Helper()

	got, err := dict.Search("test")
	if err != nil {
		t.Fatal(err)
	}

	if got != expected {
		t.Errorf("expected '%v', but got '%v'", expected, got)
	}

}

func assertStrings(t *testing.T, expected string, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%v', but got '%v'", expected, got)
	}
}

func assertError(t *testing.T, expected error, got error) {
	t.Helper()
	if expected != got {
		t.Errorf("expected error %v, but got %v", expected, got)
	}
}
