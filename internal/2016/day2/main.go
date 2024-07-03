package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

var keypad1 = [3][3]rune{
	{'1', '2', '3'},
	{'4', '5', '6'},
	{'7', '8', '9'},
}

var keys = [13]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D'}

var keypad2 = [5][5]*rune{

	{nil, nil, &keys[0], nil, nil},

	{nil, &keys[1], &keys[2], &keys[3], nil},

	{&keys[4], &keys[5], &keys[6], &keys[7], &keys[8]},

	{nil, &keys[9], &keys[10], &keys[11], nil},

	{nil, nil, &keys[12], nil, nil},
}

func main() {
	data := reader.FileToArray("data/2016/day2.txt")

	solveFirstProblem(data)
	solveSecondProblem(data)
}

func solveFirstProblem(data []string) {
	res := strings.Builder{}
	row, col := 1, 1
	for _, r := range data {
		for _, c := range r {
			switch c {
			case 'U':
				row = max(0, row-1)
			case 'L':
				col = max(0, col-1)
			case 'D':
				row = min(2, row+1)
			case 'R':
				col = min(2, col+1)
			}
		}
		res.WriteRune(keypad1[row][col])
	}

	fmt.Println(res.String())
}

func solveSecondProblem(data []string) {
	res := strings.Builder{}
	row, col := 2, 0
	for _, r := range data {
		for _, c := range r {
			switch c {
			case 'U':
				newRow := max(0, row-1)
				if keypad2[newRow][col] == nil {
					continue
				}
				row = newRow
			case 'L':
				newCol := max(0, col-1)
				if keypad2[row][newCol] == nil {
					continue
				}
				col = newCol
			case 'D':
				newRow := min(4, row+1)
				if keypad2[newRow][col] == nil {
					continue
				}
				row = newRow
			case 'R':
				newCol := min(4, col+1)
				if keypad2[row][newCol] == nil {
					continue
				}
				col = newCol
			}
		}
		res.WriteRune(*keypad2[row][col])
	}

	fmt.Println(res.String())
}
