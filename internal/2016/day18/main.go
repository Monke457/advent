package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

const ROWS = 400000

func main() {
	traps := reader.FileTo2DArray("data/2016/day18.txt")


	row := traps[0]
	for len(traps) < ROWS {
		row = buildTraps(row)
		traps = append(traps, row)
	}
	safe := countSafe(traps)
	fmt.Println("Safe tiles:", safe)
}

func countSafe(traps [][]rune) int {
	safe := 0
	for _, row := range traps {
		for _, tile := range row {
			if tile == '.' {
				safe++
			}
		}
	}
	return safe
}

func buildTraps(row []rune) []rune {
	next := []rune{}
	for i := 0; i < len(row); i++ {
		if i == 0 {
			if row[i+1] == '^' {
				next = append(next, '^')
			} else {
				next = append(next, '.')
			}
		} else if i == len(row)-1 {
			if row[i-1] == '^' {
				next = append(next, '^')
			} else {
				next = append(next, '.')
			}
		} else {
			if (row[i-1] == '^' && row[i+1] != '^') || (row[i+1] == '^' && row[i-1] != '^') {
				next = append(next, '^')
			} else {
				next = append(next, '.')
			}
		}
	}
	return next
}
