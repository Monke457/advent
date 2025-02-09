package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
)

func main() {
	data := reader.FileTo2DIntArray("data/2021/day9.txt")

	lowpoints := findLowpoints(data) 
	sum := len(lowpoints)
	for _, pos := range lowpoints {
		sum += data[pos[1]][pos[0]]
	}
	fmt.Println("Sum of risk levels:", sum)

	basins := findBasins(data, lowpoints)
	product := 1
	for i := range 3 {
		product *= len(basins[i])
	}

	fmt.Println("Product of biggest 3 basins:", product)
}

func findBasins(data [][]int, lowpoints [][2]int) []map[[2]int]bool {
	basins := []map[[2]int]bool{}
	for _, lowpoint := range lowpoints {
		basins = append(basins, findBasin(data, lowpoint))
	}
	slices.SortFunc(basins, func(a, b map[[2]int]bool) int {
		if len(a) > len(b) {
			return -1
		}
		return 1
	})
	return basins
}

func findBasin(data [][]int, start [2]int) map[[2]int]bool {
	queue := map[[2]int]int{start:0}
	visited := map[[2]int]bool{}
	curr := [2]int{}
	currVal := 0

	for len(queue) > 0 {
		queue, curr, currVal = pop(queue)
		if visited[curr] {
			continue
		}
		visited[curr] = true
		for _, n := range getNeighbours(data, curr) {
			if visited[n] {
				continue
			}
			val := data[n[1]][n[0]] 
			if val == 9 {
				continue
			}
			if val < currVal {
				continue
			}
			queue[n] = val
		}
	}
	return visited
}

func pop(queue map[[2]int]int) (map[[2]int]int, [2]int, int) {
	var node [2]int
	best := math.MaxInt
	for k, v := range queue {
		if v < best {
			best = v
			node = k
		}
	}
	delete(queue, node)
	return queue, node, best
}

func findLowpoints(data [][]int) [][2]int {
	lowpoints := [][2]int{}
	for y := range data {
		loop:
		for x, val := range data[y] {
			pos := [2]int{x, y}
			for _, n := range getNeighbours(data, pos) {
				if data[n[1]][n[0]] <= val {
					continue loop
				}
			}
			lowpoints = append(lowpoints, pos)
		}
	}
	return lowpoints
}

var dirs = [][2]int{ {-1, 0}, {1, 0}, {0, 1}, {0, -1} }

func getNeighbours(data [][]int, pos [2]int) [][2]int {
	neighbours := [][2]int{}
	for _, dir := range dirs {
		neighbour := [2]int{ pos[0] + dir[0], pos[1] + dir[1] }
		if oob(data, neighbour) {
			continue
		}
		neighbours = append(neighbours, neighbour)
	}
	return neighbours
}

func oob(data [][]int, pos [2]int) bool {
	if pos[1] < 0 || pos[1] >= len(data) {
		return true 
	}
	if pos[0] < 0 || pos[0] >= len(data[pos[1]]) {
		return true 
	}
	return false
}
