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
func Test_borderString(t *testing.T) {
	tbl := new(TextTable)
	tbl.maxWidths = []int{4, 5, 3, 2}

	expected := "+------+-------+-----+----+"

	border := tbl.borderString()
	if border != expected {
		t.Errorf("got %s(Expected %s)", border, expected)
	}

	tbl.maxWidths = []int{0}
	expected = "+--+"
	border = tbl.borderString()
	if border != expected {
		t.Errorf("got %s(Expected %s)", border, expected)
	}
}
