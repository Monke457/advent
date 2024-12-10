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

	next := val+1
	n := [2]int{y-1, x}
	e := [2]int{y, x+1}
	s := [2]int{y+1, x}
	w := [2]int{y, x-1}

	score := 0
	rating := 0
	if !oob(data, n[0], n[1]) && data[n[0]][n[1]] == next {
		s, r := evaluateTrail(data, start, n[0], n[1])
		score += s
		rating += r
	}
	if !oob(data, e[0], e[1]) && data[e[0]][e[1]] == next {
		s, r := evaluateTrail(data, start, e[0], e[1])
		score += s
		rating += r
	}
	if !oob(data, s[0], s[1]) && data[s[0]][s[1]] == next {
		s, r := evaluateTrail(data, start, s[0], s[1])
		score += s
		rating += r
	}
	if !oob(data, w[0], w[1]) && data[w[0]][w[1]] == next {
		s, r := evaluateTrail(data, start, w[0], w[1])
		score += s
		rating += r
	}
	return score, rating
}
