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

func Test_formatCellUnit(t *testing.T) {
	cell := cellUnit{content: "apple", alignment: ALIGN_RIGHT}

	expected := " apple "
	got := formatCellUnit(&cell, 5)
	if got != expected {
		t.Errorf("got '%s'(Expected '%s')", got, expected)
	}

	expected = "      apple "
	got = formatCellUnit(&cell, 10)
	if got != expected {
		t.Errorf("got '%s'(Expected '%s')", got, expected)
	}

	cellLeft := cellUnit{content: "orange", alignment: ALIGN_LEFT}
	expected = " orange "
	got = formatCellUnit(&cellLeft, 6)
	if got != expected {
		t.Errorf("got '%s'(Expected '%s')", got, expected)
	}

	expected = " orange     "
	got = formatCellUnit(&cellLeft, 10)
	if got != expected {
		t.Errorf("got '%s'(Expected '%s')", got, expected)
	}
}

func Test_generateRowString(t *testing.T) {
	tbl := TextTable{}
	tbl.maxWidths = []int{8, 5}
	cells := []*cellUnit{
		&cellUnit{content: "apple", alignment: ALIGN_RIGHT},
		&cellUnit{content: "melon", alignment: ALIGN_RIGHT},
	}

	row := tableRow{ cellUnits: cells, kind: ROW_CELLS,}
	got := tbl.generateRowString(&row)

	expected := "|    apple | melon |"
	if got != expected {
		t.Errorf("got '%s'(Expected '%s')", got, expected)
	}
}
