package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2021/day5.txt")

	points := map[[2]int]int{}
	all := map[[2]int]int{}

	for _, line := range data {
		vent, diagonal := getVentPoints(line)
		for _, point := range vent {
			all[point]++
			if !diagonal {
				points[point]++
			}
		}
	}

	first := 0
	second := 0
	for key := range all {
		if points[key] >= 2 {
			first++
		}
		if all[key] >= 2 {
			second++
		}
	}

	fmt.Println("dangerous points (no diagonals):", first)
	fmt.Println("dangerous points (diagonals):", second)
}

func getVentPoints(line string) ([][2]int, bool) {
	p1, p2, _ := strings.Cut(line, " -> ")
	x1str, y1str, _ := strings.Cut(p1, ",")
	x2str, y2str, _ := strings.Cut(p2, ",")

	x1, _ := strconv.Atoi(x1str)
	x2, _ := strconv.Atoi(x2str)
	y1, _ := strconv.Atoi(y1str)
	y2, _ := strconv.Atoi(y2str)

	if (x1 != x2 && y1 != y2) || (x1 == x2 && y1 == y2) {
		return buildDiagonal(x1, y1, x2, y2), true
	}
	return buildLine(x1, y1, x2, y2), false
}

func buildDiagonal(x1, y1, x2, y2 int) [][2]int {
	xc, yc := 1, 1
	xdiff := x2 - x1
	if xdiff < 0 {
		xdiff = -xdiff
		xc = -1
	}
	if y2 - y1 < 0 {
		yc = -1
	}

	points := [][2]int{}
	for i := range xdiff+1 {
		points = append(points, [2]int{x1+xc*i, y1+yc*i})
	} 
	return points
}

func buildLine(x1, y1, x2, y2 int) [][2]int {
	points := [][2]int{}
	if x1 > x2 {
		x1 ^= x2
		x2 ^= x1
		x1 ^= x2
	}
	if y1 > y2 {
		y1 ^= y2
		y2 ^= y1
		y1 ^= y2
	}
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			points = append(points, [2]int{i, j})
		}
	} 
	return points
}

func diagonal(x1, y1, x2, y2 int) bool {
	inline := 0
	if x1 == x2  {
		inline++
	}
	if y1 == y2 {
		inline++
	}
	return inline != 1
}
