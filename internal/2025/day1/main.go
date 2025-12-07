package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
)

func main() {
	data := reader.FileToArray("data/2025/day1.txt")

	var val, count, turns int

	var pol int = 1
	var m int = 100
	var curr int = 50

	for _, line := range data {
		val, _ = strconv.Atoi(line[1:])
		if line[0] == 'L' { val *= -1 }

		curr = curr + val
		turns = curr / m
		curr %= m

		if turns < 0 { turns *= -1 }

		if curr == 0 { 
			if turns == 0 || (pol < 0 && line[0] == 'R') || (pol > 0 && line[0] == 'L') {
				turns++ 
			}
			pol = 0
		} else if curr < 0 {
			if pol > 0 {
				turns++
			}
			pol = -1
		} else if curr > 0  {
			if pol < 0 {
				turns++
			}
			pol = 1
		}
		count += turns
	}
	fmt.Println("times at zero:", count)
}
