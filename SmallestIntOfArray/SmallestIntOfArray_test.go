package kata

import "testing"

func TestSmallestIntOfArray(t *testing.T) {
	got := SmallestIntegerFinder([]int{0, 1, 2, 3})
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SmallestIntegerFinder([]int{34, 15, 88, 2})
	want = 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SmallestIntegerFinder([]int{34, -345, -1, 100})
	want = -345
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
