package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got : %d, want: %d, given: %v ", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// want := "bob"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want: %d", got, want)
	}

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	t.Run("Sum tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 4, 5}, []int{0, 9})
		want := []int{9, 9}

		checkSums(t, got, want)
	})
	t.Run("Sum tails even when slice is empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 4})
		want := []int{0, 4}

		checkSums(t, got, want)
	})

}
