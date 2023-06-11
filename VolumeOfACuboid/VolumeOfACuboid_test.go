package kata

import "testing"

/*
DESCRIPTION:
Bob needs a fast way to calculate the volume of a cuboid with three values:
the length, width and height of the cuboid. Write a function to help Bob with this calculation.
*/

var (
	got, want float64
	cases     = []struct {
		w, l, h, v float64
	}{
		{0, 0, 0, 0},
	}
)

func TestGetVolumeOfCuboid(t *testing.T) {
	for _, val := range cases {
		got = GetVolumeOfCuboid(val.l, val.w, val.h)
		want = val.v
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	}
}
