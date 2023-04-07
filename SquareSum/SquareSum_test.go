package kata

import "testing"

// SquareSum([]int{1,2}) // Equal(5)
// SquareSum([]int{0,3,4,5}) // Equal(50)
// SquareSum([]int{}) // Equal(0)

func TestSquareSum(t *testing.T) {
	got := SquareSum([]int{1, 2})
	want := 5
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}

	got = SquareSum([]int{0, 3, 4, 5})
	want = 50
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}

	got = SquareSum([]int{})
	want = 0
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
