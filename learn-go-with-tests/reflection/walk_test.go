package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	t.Run("care about the order of elements", func(t *testing.T) {

		cases := []struct {
			Name     string
			Input    interface{}
			Expected []string
		}{
			{
				"Struct with one field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct with two fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"Struct with non-string fields",
				struct {
					Name string
					Age  int
				}{"Chris", 18},
				[]string{"Chris"},
			}, {
				"Nested struct",
				Person{"Chris", Profile{18, "London"}},
				[]string{"Chris", "London"},
			}, {
				"Pointer to something",
				&Person{"Chris", Profile{18, "London"}},
				[]string{"Chris", "London"},
			}, {
				"Slices",
				[]Profile{
					{12, "A"},
					{13, "B"},
				},
				[]string{"A", "B"},
			}, {
				"Arrays",
				[2]Profile{
					{12, "A"},
					{13, "B"},
				},
				[]string{"A", "B"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				Walk(test.Input, func(s string) {
					got = append(got, s)
				})
				if !reflect.DeepEqual(test.Expected, got) {
					t.Errorf("expected %v, got %v", test.Expected, got)
				}
			})
		}
	})
	t.Run("care about the order with map", func(t *testing.T) {
		m := map[string]string{
			"A": "AA",
			"B": "BB",
		}

		var got []string
		Walk(m, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "AA")
		assertContains(t, got, "BB")
	})
}

func assertContains(t *testing.T, heystack []string, needle string) {
	contains := false
	for _, v := range heystack {
		if v == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", heystack, needle)
	}

}
