package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	example := Add(1, 2)
	fmt.Println(example)
	// Output:
	// 3
}

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	excepted := 3
	if got != excepted {
		t.Errorf("excepted %d, got %d", excepted, got)
	}
}