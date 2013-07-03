package texttable

import (
	"testing"
)


//
// Private Function/Methods
//

func Test_calcMaxHeight(t *testing.T) {
	input := []string{
		"hello", "apple\nmelon\norange", "1\n2",
	}

	got := calcMaxHeight(input)
	if got != 3 {
		t.Errorf("calcMaxHeight(%s) != 3(got=%d)", input, got)
	}
}

func Test_divideByNewLine(t *testing.T) {
	input := "apple\nmelon\norange"

	got := divideByNewLine(input)
	if len(got) != 3 {
		t.Errorf("length of retval != 3(got=%d)", input, len(got))
	}

	if !(got[0] == "apple" && got [1] == "melon" && got [2] == "orange") {
		t.Errorf("expected apple, melon, orange(got %s, %s %s)",
			got[0], got[1], got[2])
	}
}
