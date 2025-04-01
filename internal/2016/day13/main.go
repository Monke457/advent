package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
)

func main() {
	data := reader.FileToIntArray("data/2016/day13.txt")

	dest := [2]int{31, 39}
	path := findPath(dest, data[0])
	positions := findPositionsWithLimit(data[0], 50)

	fmt.Println("Length of path:", len(path)-1)
	fmt.Println("Possible locations within 50 steps:", positions)
}

func findPositionsWithLimit(num, limit int) int {
	start := [2]int{1,1}
	queue := map[[2]int]int{start:0} 
	costs := map[[2]int]int{start:0}
	visited := map[[2]int]bool{}

	var node [2]int
	for len(queue) > 0 {
		node = pop(&queue)
		if visited[node] {
			continue
		}
		visited[node] = true

		for _, next := range GetNeighbours(node, num) {
			if visited[next] {
				continue
			}
			cost := costs[node] + 1
			if cost > limit {
				continue
			}
			if c, ok := costs[next]; !ok || cost < c {
				costs[next] = cost
				queue[next] = cost 
			}
		}
	}
	return len(visited)
}

func findPath(dest [2]int, num int) [][2]int {
	start := [2]int{1,1}
	queue := map[[2]int]int{start:0} 
	costs := map[[2]int]int{start:0}
	predecessors := map[[2]int]*[2]int{start:nil}
	visited := map[[2]int]bool{}

	var node [2]int
	for len(queue) > 0 {
		node = pop(&queue)
		if node == dest {
			break
		}
		if visited[node] {
			continue
		}
		visited[node] = true

		for _, next := range GetNeighbours(node, num) {
			if visited[next] {
				continue
			}
			cost := costs[node] + 1
			if c, ok := costs[next]; !ok || cost < c {
				costs[next] = cost
				queue[next] = cost + heuristic(next, dest)
				temp := [2]int{node[0], node[1]}
				predecessors[next] = &temp
			}
		}
	}
	return makePath(predecessors, dest)
}

func heuristic(pos, dest [2]int) int {
	yDiff := math.Abs(float64(dest[0] - pos[0]))
	xDiff := math.Abs(float64(dest[1] - pos[1]))
	return int(yDiff + xDiff)
}

var dirs = [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
var cache = map[[2]int]bool{}

func GetNeighbours(pos [2]int, num int) [][2]int {
	neighbours := [][2]int{}
	for _, dir := range dirs {
		y, x := pos[0] + dir[0], pos[1] + dir[1]
		if y < 0 || x < 0 {
			continue
		}

		neighbour := [2]int{y, x}
		if v, ok := cache[neighbour]; ok {
			if v {
				neighbours = append(neighbours, neighbour)
			}
		} else if isOpen(y, x, num) {
			neighbours = append(neighbours, neighbour)
			cache[neighbour] = true
		} else {
			cache[neighbour] = false
		}
	}
	return neighbours
}

func isOpen(y, x, num int) bool {
	val := y*y + 3*y + 2*y*x + x + x*x + num
	bin := fmt.Sprintf("%b", val)
	count := 0
	for _, bit := range bin {
		if bit == '1' {
			count++
		}
	}
	return count % 2 == 0
}

func pop(queue *map[[2]int]int) [2]int {
	var key [2]int
	val := math.MaxInt
	for k, v := range *queue {
		if v < val {
			key = k
			val = v
		}
	}
	delete(*queue, key)
	return key
}

func makePath(predecessors map[[2]int]*[2]int, dest [2]int) [][2]int {
	path := [][2]int{dest}
	next := predecessors[dest]
	for next != nil {
		path = append(path, *next)
		next = predecessors[*next]
	}
	slices.Reverse(path)
	return path
}
