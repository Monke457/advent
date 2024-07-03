package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2020/day3.txt")


	first := countTrees([2]int{1,3}, data)
	fmt.Println("First:", first)

	slopes := [][2]int{{1,1}, {1,5}, {1,7}, {2,1}}
	product := first
	for _, slope := range slopes {
		product *= countTrees(slope, data)
	}
	fmt.Println("Second:", product)
}

func countTrees(s [2]int, data []string) int {
	var row, col, count int 
	cols := len(data[0])
	for {
		row = (row + s[0])
		if row >= len(data) {
			break
		}
		col = (col + s[1]) % cols
	 
		if data[row][col] == '#' {
			count++
		}
	}
	return count
}
