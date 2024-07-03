package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArray("data/2017/day5.txt")

	steps := walk(data)
	fmt.Println("First:", steps)
}

func walk(data []int) int {
	steps := 0
	pos := 0
	for pos >= 0 && pos < len(data) {
		tmp := pos
		pos += data[pos]
		if data[tmp] >= 3 {
			data[tmp]--
		} else {
			data[tmp]++
		}
		steps++
	}
	return steps
}
