package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileTo2DIntArray("data/2024/day10.txt")

	sumScore := 0
	sumRating := 0
	
	for i := range data {
		for j := range data[i]{
			if data[i][j] == 0 {
				s, r := evaluateTrail(data, [2]int{i, j}, i, j)
				sumScore += s
				sumRating += r
			}
		}
	}

	fmt.Println("score sum:", sumScore)
	fmt.Println("rating sum:", sumRating)
}

func oob(data [][]int, y, x int) bool {
	if y < 0 || y >= len(data) {
		return true
	}
	if x < 0 || x >= len(data[y]) {
		return true
	}
	return false
}

var cache = map[[2]int][][2]int{}

func evaluateTrail(data [][]int, start [2]int, y, x int) (int, int) {
	val := data[y][x]
	if val == 0 {
		cache[start] = [][2]int{}
	}
	if val == 9 { 
		for _, ends := range cache[start] {
			if ends[0] == y && ends[1] == x {
				return 0, 1
			}
		}
		cache[start] = append(cache[start], [2]int{y, x})
		return 1, 1
	}

	dirs := [][2]int{
		{y-1, x},
		{y, x+1},
		{y+1, x},
		{y, x-1},
	}

	score := 0
	rating := 0
	for _, d := range dirs {
		if !oob(data, d[0], d[1]) && data[d[0]][d[1]] == val+1 {
			s, r := evaluateTrail(data, start, d[0], d[1])
			score += s
			rating += r
		}
	}
	return score, rating
}
