package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum in array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		expected := 15
		if expected != got {
			t.Errorf("expected: %d, but got: %d, given: %v", expected, got, nums)
		}
	})

	t.Run("sum in slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		got := Sum(slice)
		expected := 15
		if expected != got {
			t.Errorf("expected: %d, but got: %d, given: %v", expected, got , slice)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {

	assertSum := func(expected []int, sum []int) {
		if !reflect.DeepEqual(expected, sum) {
			t.Errorf("expected %v, got %v", expected, sum)
		}
	}

	t.Run("valid array", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		expected := []int{2, 4}
		assertSum(expected, got)
	})

	t.Run("empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4})
		expected := []int{0, 4}
		assertSum(expected, got)
	})
}
