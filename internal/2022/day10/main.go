package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2022/day10.txt")

	idx := 0
	x := 1
	executing := false
	queue := 0
	crt := [6][]rune{}

	for cycle := 1; cycle <= 240; cycle++ {
		row := (cycle-1)/40
		col := (cycle-1)%40
		var val rune

		if col == x || col == x-1 || col == x+1 {
			val = '#'
		} else {
			val = '.'
		}

		crt[row] = append(crt[row], val)

		if executing {
			executing = false
			x += queue
			queue = 0
			continue
		}

		op, valstr, _ := strings.Cut(data[idx], " ")
		if op == "addx" {
			val, _ := strconv.Atoi(valstr)
			queue = val
			executing = true
		}
		idx++
	}

	for _, row := range crt {
		fmt.Printf("%s\n", string(row))
	}
}
