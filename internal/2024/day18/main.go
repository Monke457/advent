package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const SIZE int = 71
const BYTES int = 1024

func main() {
	data := reader.FileToArray("data/2024/day18.txt")

	muhmap := parseMap(data[:BYTES])
	path, _ := findPath(muhmap)
	steps := len(path)

	i := 1
	for {
		muhmap = dropByte(muhmap, data[BYTES+i])
		if _, found := findPath(muhmap); !found {
			break
		}
		i++
	}

	fmt.Println("steps needed:", steps)
	//takes a few seconds...
	fmt.Println("Coordinates of the blocking byte:", data[BYTES+i])
}

func findPath(m [][]rune) ([][2]int, bool) {
	src := [2]int{0, 0}
	dest := [2]int{SIZE-1, SIZE-1}
	visited := map[[2]int]bool{}
	costs := map[[2]int]int{src:0}
	queue := map[[2]int]int{src:0}
	predecessors := map[[2]int]*[2]int{src:nil}

	var next [2]int
	for {
		if len(queue) == 0 {
			break
		}
		queue, next = pop(queue)
		if next == dest {
			break
		}
		if visited[next] {
			continue
		}
		visited[next] = true
		for _, n := range getNeighbours(next) {
			if m[n[0]][n[1]] == '#' {
				continue
			}
			curr_cost := costs[next] + 1
			if prev, ok := costs[n]; !ok || curr_cost < prev {
				costs[n] = curr_cost
				queue[n] = costs[n] + heuristic(next, n)
				pred := [2]int{next[0], next[1]}
				predecessors[n] = &pred
			}
		}
	}

	path := [][2]int{dest}
	if _, ok := predecessors[dest]; !ok {
		return path, false
	}

	pred := predecessors[dest]
	for pred != nil {
		path = append(path, *pred)
		pred = predecessors[*pred]
	}

	slices.Reverse(path)
	return path, true
}

func pop(queue map[[2]int]int) (map[[2]int]int, [2]int) {
	key, val := [2]int{}, math.MaxInt 
	for k, v := range queue {
		if v < val {
			val = v
			key = k
		}
	}
	delete(queue, key)
	return queue, key 
}

func heuristic(pos, dest [2]int) int {
	return int(math.Abs(float64(dest[0] - pos[0])) + math.Abs(float64(dest[1] - pos[1])))
}

var dirs = [4][2]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func getNeighbours(pos [2]int) [][2]int {
	n := [][2]int{}
	for _, dir := range dirs {
		next := [2]int{pos[0]+dir[0], pos[1]+dir[1]}
		if inRange(next) {
			n = append(n, next)
		}
	}
	return n
}

func inRange(pos [2]int) bool {
	if pos[0] < 0 || pos[0] >= SIZE {
		return false
	}
	if pos[1] < 0 || pos[1] >= SIZE {
		return false
	}
	return true
}

func dropByte(m [][]rune, str string) [][]rune {
	x, y := parseCoords(str)
	m[y][x] = '#'
	return m
}

func parseMap(data []string) [][]rune{
	result := make([][]rune, SIZE)
	for i := range SIZE {
		result[i] = make([]rune, SIZE)
		for j := range SIZE {
			result[i][j] = '.'
		}
	}
	for _, line := range data {
		x, y := parseCoords(line)
		result[y][x] = '#'
	}
	return result
} 

func parseCoords(str string) (int, int) {
	xStr, yStr, _ := strings.Cut(str, ",")
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	return x, y
}

func drawmap(m [][]rune, p [][2]int) {
	mwp := make([][]rune, len(m))
	copy(mwp, m)
	for i := range m {
		mwp[i] = make([]rune, len(m[i]))
		copy(mwp[i], m[i])
	}
	for _, coord := range p {
		mwp[coord[0]][coord[1]] = 'O'
	}

	for _, row := range mwp {
		fmt.Printf("%c\n", row)
	}
}
