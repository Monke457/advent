package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
)

func main() {
	data := reader.FileTo2DArray("data/2024/day20.txt")

	path := getPath(data)
	limit := 20

	cheats := findCheats(path, limit)

	savedTimes := map[int]int{}
	for start, ends := range cheats {
		for end := range ends {
			saved := end - start
			dist := math.Abs(float64(path[start][0] - path[end][0])) + math.Abs(float64(path[start][1] - path[end][1])) 
			savedTimes[saved-int(dist)]++
		}
	}

	results := 0
	target := 100
	for k, v := range savedTimes {
		if k >= target {
			results += v
		}
	}

	fmt.Println("There are", results, "cheats that save at least", target, "picoseconds")
}

func getPath(maze [][]rune) [][2]int {
	pos := findRune(maze, 'S')
	end := findRune(maze, 'E')
	path := [][2]int{pos}
	step := 0
	for pos != end {
		pos = getNext(maze, path[max(0, step-1)], pos)
		path = append(path, pos)
		step++
	}
	return path
}

var dirs = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
func getNext(maze [][]rune, prev, curr [2]int) [2]int {
	for _, dir := range dirs {
		x, y := curr[0] + dir[0], curr[1] + dir[1]
		if maze[y][x] == '#' {
			continue
		}
		if prev[0] == x && prev[1] == y {
			continue
		}
		return [2]int{x, y}
	}
	panic(fmt.Errorf("Couldn't find the next step at position %d", curr))
}


func findCheats(path [][2]int, limit int) map[int]map[int]bool {
	cheats := map[int]map[int]bool{}
	for i, node := range path {

		endings := findEndPoints(path, node, limit)
		if len(endings) == 0 {
			continue
		}

		if _, ok := cheats[i]; !ok {
			cheats[i] = map[int]bool{}
		}
		for _, ending := range endings {
			cheats[i][ending] = true
		}
	}
	return cheats
}

func getIndex(path [][2]int, pos [2]int) int {
	for i, p := range path {
		if p == pos {
			return i
		}
	}
	panic(fmt.Errorf("position %d could not be found in path %v", pos, path))
}

func findEndPoints(path [][2]int, pos [2]int, limit int) []int {
	endings := []int{}
	idx := getIndex(path, pos)

	for i, node := range path[idx+1:] {
		dist := math.Abs(float64(pos[0] - node[0])) + math.Abs(float64(pos[1] - node[1]))
		if i+1 <= int(dist) || int(dist) > limit {
			continue
		}
		endings = append(endings, idx + i + 1)
	}
	return endings
}

func draw(maze [][]rune, path [][2]int) {
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == 'S' || maze[y][x] == 'E' {
				fmt.Printf("%c", maze[y][x])
				continue
			}
			if slices.Contains(path, [2]int{x, y}) {
				fmt.Printf("O")
				continue
			}
			fmt.Printf("%c", maze[y][x])
		}
		fmt.Println()
	}
}

func findRune(data [][]rune, r rune) [2]int {
	for y := range data {
		for x := range data[y] {
			if data[y][x] == r {
				return [2]int{x, y}
			}
		}
	}
	panic(fmt.Errorf("Could not find rune %c in data", r))
}
