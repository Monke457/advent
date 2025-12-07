package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileTo2DArray("data/2025/day4.txt")

	count := 0
	flagRemove := map[int][]int{}
	for {
		for i := range data {
			for j, r := range data[i] {
				if r != '@' { continue }
				if countNeighbours(data, i, j) >= 4 { continue }
				count++
				flagRemove[i] = append(flagRemove[i], j)
			}
		}
		if len(flagRemove) == 0 { break }
		data = removeRolls(data, flagRemove)
		clear(flagRemove)
	}
	fmt.Println(count)
}

func removeRolls(data [][]rune, flagRemove map[int][]int) [][]rune {
	data_new := [][]rune{}
	for i := range data {
		if remove, ok := flagRemove[i]; ok {
			for _, r := range remove {
				data[i][r] = '.'
			}
		}
		data_new = append(data_new, data[i])
	}
	return data_new
}

var dirs = [][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}

func countNeighbours(data [][]rune, x, y int) int {
	count := 0
	var x_new, y_new int
	for _, dir := range dirs {
		x_new = x + dir[0]
		y_new = y + dir[1]
		if oob(data, x_new, y_new) { continue }
		if data[x_new][y_new] == '@' { count++ }
	} 
	return count
}

func oob(data [][]rune, x, y int) bool {
	if x < 0 || y < 0 { return true }
	if x >= len(data) { return true }
	return y >= len(data[x])
}
