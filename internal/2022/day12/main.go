package main

import (
	"advent/internal/pkg/grid"
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"

	pg "github.com/Monke457/printge"
)

func main() {
	data := reader.FileTo2DArray("data/2022/day12.txt")

	start := findRune(data, 'S')
	end := findRune(data, 'E')

	data[start[0]][start[1]] = 'a'
	data[end[0]][end[1]] = 'z'

	var best [][2]int
	l := math.MaxInt

	for _, start := range findAllRunes(data, 'a') {
		route, ok := getShortestRoute(data, start, end)
		if !ok {
			continue
		}
		length := len(route)-1
		if length < l {
			l = length
			best = route
		}
	}
	drawRoute(data, best)
	fmt.Println("Shortest route from lowest point:", l)
}

func drawRoute(data [][]rune, route [][2]int) {
	for y := range data {
		for x := range data[y] {
			if slices.Contains(route, [2]int{y, x}) {
				pg.Print(fmt.Sprintf("%c", data[y][x]), pg.WHITE)
			} else {
				fmt.Printf("%c", data[y][x])
			}
		}
		fmt.Println()
	}
}

func getShortestRoute(data [][]rune, start, end [2]int) ([][2]int, bool) {
	queue := map[[2]int]int{start:0}
	visited := map[[2]int]bool{}
	costs := map[[2]int]int{start:0}
	path := map[[2]int]*[2]int{start:nil}

	var current [2]int
	for len(queue) > 0 {
		current = pop(&queue)
		if current == end {
			break
		}
		if visited[current] {
			continue
		}
		visited[current] = true


		for _, n := range grid.GetNeighboursCont(data, current) {
			cost, ok := getCost(data, current, n)
			if !ok {
				continue
			}
			cost += costs[current] + heuristic(n, end)
			if _, ok := costs[n]; !ok || cost < costs[n] {
				costs[n] = cost
				queue[n] = cost 
				temp := [2]int{current[0], current[1]}
				path[n] = &temp
			}
		}
	}

	return makeRoute(path, start, end)
}

func heuristic(a, b [2]int) int {
	return int(math.Abs(float64(a[0] - b[0])) + math.Abs(float64(a[1] - b[1])))
}

func getCost(data [][]rune, a, b [2]int) (int, bool) {
	rA := data[a[0]][a[1]]
	rB := data[b[0]][b[1]]
	cost := int(rB) - int(rA)
	if cost > 1 {
		return -1, false
	}
	return -cost, true
}

func makeRoute(path map[[2]int]*[2]int, start, dest [2]int) ([][2]int, bool) {
	route := [][2]int{dest}
	next := path[dest]
	for next != nil {
		route = append(route, *next)
		next = path[*next]
	}
	slices.Reverse(route)
	if route[0] != start {
		return route, false
	}
	return route, true
}

func pop(queue *map[[2]int]int) [2]int {
	var node [2]int
	score := math.MaxInt
	for key, val := range *queue {
		if val < score {
			score = val
			node = key
		}
	}
	delete(*queue, node)
	return node
}

func findRune(data [][]rune, r rune) [2]int {
	for y := range data {
		for x := range data[y] {
			if data[y][x] == r {
				return [2]int{y, x}
			}
		}
	}
	panic(fmt.Errorf("could not find position '%c' in data", + r))
}

func findAllRunes(data [][]rune, r rune) [][2]int {
	positions := [][2]int{}
	for y := range data {
		for x := range data[y] {
			if data[y][x] == r {
				positions = append(positions, [2]int{y, x})
			}
		}
	}
	return positions
}

