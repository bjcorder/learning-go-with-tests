package sum

import "testing"

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
