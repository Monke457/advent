package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

func main() {
	data := reader.FileTo2DArray("data/2024/day8.txt")
	draw(data)

	antennae := parseAntennae(data)

	antinodes := map[[2]int]bool{}
	for _, positions := range antennae { 
		for i := 0; i < len(positions); i++ {
			for j := i+1; j < len(positions); j++ {
				for key := range getAntinodes(data, positions[i], positions[j]) {
					antinodes[key] = true
				}
			}
		}
	}

	draw(data)
	fmt.Println("Antinodes:", len(antinodes))
}

func getAntinodes(data [][]rune, a, b [2]int) map[[2]int]bool {
	nodes := map[[2]int]bool{}
	yDiff := b[0] - a[0]
	xDiff := b[1] - a[1]
	c := float64(xDiff) / float64(yDiff)
	m := float64(a[1]) - float64(a[0]) * c

	for y := 0; y < len(data); y++ {
		res := c * float64(y) + m
		if math.Abs(res - math.Round(res)) > 0.00001 {
			continue
		}
		x := int(math.Round(res))
		if inBounds(data, y, x) {
			nodes[[2]int{y, x}] = true
			if data[y][x] == '.' {
				data[y][x] = '#'
			}
		}
	}
	return nodes
}

func inBounds(data [][]rune, y, x int) bool {
	if y < 0 || y >= len(data) {
		return false
	}
	if x < 0 || x >= len(data[y]) {
		return false
	}
	return true 
}

func parseAntennae(data [][]rune) map[rune][][2]int {
	result := map[rune][][2]int{}

	for y, row := range data {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			if _, ok := result[cell]; !ok {
				result[cell] = [][2]int{}
			}
			result[cell] = append(result[cell], [2]int{y, x})
		}
	}

	return result
}

func draw(data [][]rune) {
	fmt.Println()
	for _, row := range data {
		fmt.Printf("%c\n", row)
	}
	fmt.Println()
}
