package main

import (
	"advent/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2015/day1.txt")

	floor := 0
	position := 0
	for i, r := range data {
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		}
		if floor < 0 && position == 0 {
			position = i + 1
		}
	}
	fmt.Println("First problem:", floor)
	fmt.Println("Second problem:", position)
}
