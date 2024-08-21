package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		sliceNumbers := []int{1, 2, 3}

		got := Sum(sliceNumbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d', given %v", got, want, sliceNumbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("sum of each slice provided", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{1, 2, 3})
		want := []int{3, 9, 6}
		//"

		/*
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got '%d' want '%d'", got, want)
			}
		*/
		checkSumsMatch(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if slices.Compare(got, want) != 0 {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}
	t.Run("sum tails of all slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{1, 2, 3})
		want := []int{2, 9, 5}

		checkSums(t, got, want)
	})
	t.Run("sum tails of all slices, with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 0, 6, 7}, []int{0, 9, 7, 3}, []int{})
		want := []int{15, 19, 0}
		checkSums(t, got, want)
	})
}

func checkSumsMatch(t testing.TB, got []int, want []int) {
	t.Helper()
	if slices.Compare(got, want) != 0 {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
