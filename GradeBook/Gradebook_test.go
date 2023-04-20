package kata

import "testing"

/*
DESCRIPTION:
Grade book
Complete the function so that it finds the average of the three scores passed to it and returns the letter value associated with that grade.

Numerical Score	Letter Grade
90 <= score <= 100  'A'
80 <= score < 90    'B'
70 <= score < 80    'C'
60 <= score < 70    'D'
0 <= score < 60	    'F'
Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.
*/

func TestGetGrade(t *testing.T) {
	var (
		got, want rune
		cases     = []struct {
			grades [3][3]int
			grade  rune
		}{
			{[...][3]int{{95, 90, 93}, {100, 85, 96}, {92, 93, 94}}, 'A'},
			{[...][3]int{{70, 70, 100}, {82, 85, 87}, {84, 79, 85}}, 'B'}, // {89, 89, 90}}, 'B'}, // haven't figured out how to have a variable number of cases here
			{[...][3]int{{70, 70, 70}, {75, 70, 79}, {60, 82, 76}}, 'C'},
			{[...][3]int{{65, 70, 59}, {66, 62, 68}, {58, 62, 70}}, 'D'},
			{[...][3]int{{44, 55, 52}, {48, 55, 52}, {58, 59, 60}}, 'F'},
		}
	)
	for _, val := range cases {
		for _, set := range val.grades {
			got = GetGrade(set[0], set[1], set[2])
			want = val.grade
			if got != want {
				t.Errorf("got %c, want %c", got, want)
			}
		}
	}
}
