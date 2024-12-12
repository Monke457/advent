package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

func main() {
	data := reader.FileTo2DArray("data/2024/day12.txt")

	plots := parsePlots(data)

	total := 0
	discount := 0

	for _, plot := range plots {
		for _, p := range plot {
			area := len(p)
			total += getPerimeter(p) * area
			discount += getSides(p) * area
		}
	}

	fmt.Println("Total cost of fence:", total)
	fmt.Println("Discount cost of fence:", discount)
}

func parsePlots(raw [][]rune) map[rune][][][2]int {
	data := make([][]rune, len(raw))
	for i := range raw {
		row := make([]rune, len(raw[i]))
		copy(row, raw[i])
		data[i] = row
	}

	plots := map[rune][][][2]int{}
	var plot [][2]int

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == ' ' {
				continue
			}
			r := data[i][j]
			plot, data = getPlot(data, i, j, r)
			if _, ok := plots[r]; !ok {
				plots[r] = [][][2]int{}
			}
			plots[r] = append(plots[r], plot)
		}
	}
	return plots
}

func getSides(plot [][2]int) int {
	sides := [4]map[int][]int{{},{},{},{}}

	for _, pos := range plot {
		for i, d := range dirs {
			if contains(plot, pos[0] + d[0], pos[1] + d[1]) {
				continue
			}
			k, v := 0, 0
			switch i {
				case 0: k, v = pos[0], pos[1]
				case 1: k, v = pos[0]+1, pos[1]
				case 2: k, v = pos[1], pos[0]
				case 3: k, v = pos[1]+1, pos[0]
				default: continue
			}
			if _, ok := sides[i][k]; !ok {
				sides[i][k] = []int{}
			}
			sides[i][k] = append(sides[i][k], v)
		}
	}

	s := 0
	for i := range sides {
		s += len(sides[i])
		for _, side := range sides[i] {
			slices.Sort(side)
			for i := 0; i < len(side)-1; i++ {
				if side[i+1] - side[i] > 1 {
					s++
				}
			}
		}
	}

	return s
}

func getPerimeter(plot [][2]int) int {
	p := 0
	for _, pos := range plot {
		for _, d := range dirs {
			if !contains(plot, pos[0] + d[0], pos[1] + d[1]) {
				p++
			}
		}
	}
	return p
}

func contains(coords [][2]int, y, x int) bool {
	for _, coord := range coords {
		if coord[0] == y && coord[1] == x {
			return true
		}
	}
	return false
}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func getPlot(data [][]rune, y, x int, r rune) ([][2]int, [][]rune) {
	plot := [][2]int{{0, 0}}
	data[y][x] = ' '
	end := false

	for {
		if end {
			break
		}
		end = true
		for _, d := range dirs {
			for _, coords := range plot {
				newy, newx := coords[0] + y + d[0], coords[1] + x + d[1]
				if oob(data, newy, newx) {
					continue
				}
				if data[newy][newx] != r {
					continue
				}
				plot = append(plot, [2]int{newy - y, newx - x})
				data[newy][newx] = ' '
				end = false
			}
		}
	}
	return plot, data
}

func oob(data [][]rune, y, x int) bool {
	if y < 0 || y >= len(data) {
		return true
	}
	if x < 0 || x >= len(data[y]) {
		return true
	}
	return false
}
