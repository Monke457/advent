package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

type pattern struct {
	cells [][]rune
}

func main() {
	lines := reader.FileToArray("data/day13.txt")
	patterns := parsePatterns(lines)
	sum := 0
	for _, p := range patterns {
		fmt.Println()
		for _, l := range p.cells {
			fmt.Printf("%c\n", l)
		}
		sum += p.CountMirrored()
	}
	fmt.Println(sum)
}

func parsePatterns(lines []string) (res []pattern) {
	cells := [][]rune{}
	row := 0
	for i, l := range lines {
		if i+1 == len(lines) || l == "" {
			res = append(res, pattern{cells})
			cells = [][]rune{}
			row = 0
			continue
		}
		cells = append(cells, []rune{})
		for _, r := range l {
			cells[row] = append(cells[row], r)
		}
		row++
	}
	return
}

func recurseMirror(cells [][]rune, col int) *int {
	if col+1 == len(cells[0]) {
		return nil
	}
	for n, row := range cells {
		off := 0
		for col-off >= 0 && col+1+off < len(row) {
			if row[col-off] != row[col+1+off] {
				return nil
			}
			off++
		}
		if n+1 == len(cells) {
			return &col
		}
		return recurseMirror(cells[n+1:], col)
	}
	return &col
}

func verticalMirrored(cells [][]rune) *int {
	for i := 0; i < len(cells[0]); i++ {
		col := recurseMirror(cells, i)
		if col != nil {
			return col
		}
	}
	return nil
}

func horizontalMirrored(cells [][]rune) *int {
loop:
	for i := 0; i < len(cells); i++ {
		c := 0
		for i-c >= 0 && i+1+c < len(cells) {
			if !sliceEqual(cells[i-c], cells[i+1+c]) {
				continue loop
			}
			c++
		}
		return &i
	}
	return nil
}

func (p pattern) CountMirrored() int {
	v := verticalMirrored(p.cells)

	if v != nil {
		fmt.Println("vertical mirror at column", *v)
		return *v + 1
	}

	h := horizontalMirrored(p.cells)

	if h != nil {
		fmt.Println("horizontal mirror at row", *h)
		return (*h + 1) * 100
	}
	panic("no mirror found")
}

func sliceEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
