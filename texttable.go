package texttable

import (
	"fmt"
	"errors"
	"strings"
)

type cellAlignment int

const (
	ALIGN_LEFT cellAlignment = iota
	ALIGN_RIGHT
)

type tableCell struct {
	content string
	alignment cellAlignment
}

type tableRower interface{}

type tableRow struct {
	cells []tableCell
}

type tableLine struct {}

type TextTable struct {
	header []tableRow
	rows   []tableRower // tableRow or tableLine
	width  int
	border bool
}

func New(headers []string) (*TextTable, error) {
	table := new(TextTable)

	err := table.SetHeader(headers)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (t *TextTable) SetHeader(headers []string) error {
	if len(headers) == 0 {
		return errors.New("No header elements")
	}
	return nil
}

func divideByNewLine(str string) []string {
	return strings.Split(str, "\n")
}

func adjustColumns(colLines [][]string, max int) [][]string {
	adjustedCols := make([][]string, len(colLines))

	for _, lines := range colLines {
		padding := max - len(lines)
		fmt.Printf("Add %d padding line\n", padding)
		for i := 0; i < padding; i++ {
			lines = append(lines, "")
		}
		adjustedCols = append(adjustedCols, lines)
	}

	return adjustedCols
}

func (t *TextTable) divideRowsByLine(strs []string) []tableRower {
//	maxHeight := calcMaxHeight(strs)
//	strLines := make([][maxHeight]string, len(strs))
//
//	for i, str := range strs {
//		strLines[i] = divideByNewLine(str)
//	}
//
//	adjusted := adjustColumns(strLines, maxHeight)

	return nil
}

func calcMaxHeight(strs []string) int {
	max := -1

	for _, str := range strs {
		lines := strings.Split(str, "\n")
		height := len(lines)
		if height > max {
			max = height
		}
	}

	return max
}
