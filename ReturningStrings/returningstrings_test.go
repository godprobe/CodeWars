package kata

/*
Make a function that will return a greeting statement that uses an input; your program should return, "Hello, <name> how are you doing today?".

[Make sure you type the exact thing I wrote or the program may not execute properly]
*/

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("<name>")
	want := "Hello, <name> how are you doing today?"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	got = Greet("NaN")
	want = "Hello, NaN how are you doing today?"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	got = Greet("undefined")
	want = "Hello, undefined how are you doing today?"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	got = Greet("Roger the Shrubber")
	want = "Hello, Roger the Shrubber how are you doing today?"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
