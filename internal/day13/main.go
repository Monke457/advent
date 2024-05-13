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
	sumFirst := 0
	sumSecond := 0
	for _, p := range patterns {
		//first problem
		sumFirst += p.CountMirrored(false)
		//second problem
		sumSecond += p.CountMirrored(true)
	}
	fmt.Printf("first: %d \nSecond: %d\n", sumFirst, sumSecond)
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

func recurseMirrorSmudged(cells [][]rune, s bool, col int) *int {
	if col+1 == len(cells[0]) {
		return nil
	}

	for n, row := range cells {
		off := 0
		for col-off >= 0 && col+1+off < len(row) {
			if row[col-off] != row[col+1+off] {
				if s {
					return nil
				}
				s = true
			}
			off++
		}
		if n+1 == len(cells) {
			break
		}
		return recurseMirrorSmudged(cells[n+1:], s, col)
	}
	if s {
		return &col
	}
	return nil
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

func verticalMirrored(cells [][]rune, s bool) *int {
	for i := 0; i < len(cells[0]); i++ {
		var col *int
		if s {
			col = recurseMirrorSmudged(cells, false, i)
		} else {
			col = recurseMirror(cells, i)
		}
		if col != nil {
			return col
		}
	}
	return nil
}

func horizontalMirroredSmudged(cells [][]rune) *int {
loop:
	for i := 0; i < len(cells); i++ {
		smudge := false
		c := 0
		for i-c >= 0 && i+1+c < len(cells) {
			smudges := countSmudges(cells[i-c], cells[i+1+c])
			if smudges > 1 {
				continue loop
			}
			if smudges == 1 {
				if smudge {
					continue loop
				} else {
					smudge = true
				}
			}
			c++
		}
		if smudge {
			return &i
		}
	}
	return nil
}

func horizontalMirrored(cells [][]rune) *int {
loop:
	for i := 0; i < len(cells); i++ {
		c := 0
		for i-c >= 0 && i+1+c < len(cells) {
			equal := sliceEqual(cells[i-c], cells[i+1+c])
			if !equal {
				continue loop
			}
			c++
		}
		return &i
	}
	return nil
}

func (p pattern) CountMirrored(s bool) int {
	v := verticalMirrored(p.cells, s)
	if v != nil {
		return *v + 1
	}

	var h *int
	if s {
		h = horizontalMirroredSmudged(p.cells)
	} else {
		h = horizontalMirrored(p.cells)
	}

	if h != nil {
		return (*h + 1) * 100
	}
	panic("no mirror found")
}

func countSmudges(a, b []rune) int {
	if len(a) != len(b) {
		panic("arrays should be the same length")
	}
	smudges := 0
	for i := range a {
		if a[i] != b[i] {
			smudges++
		}
	}
	return smudges
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
