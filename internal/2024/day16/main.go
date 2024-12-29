package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
)

type direction int
const (
	NORTH direction = iota
	SOUTH
	EAST
	WEST
)

func main() {
	data := reader.FileTo2DArray("data/2024/day16.txt")

	path := walk(data, EAST)
	sum := sumPath(path, EAST)
	
	draw(data, path)
	fmt.Println("Best score:", sum)
}

func draw(data [][]rune, path [][2]int) {
	for y := range data {
		for x := range data[y] {
			if slices.Contains(path, [2]int{x, y}) {
				fmt.Printf("+")
			} else {
				fmt.Printf("%c", data[y][x])
			}
		}
		fmt.Println()
	}
}

func walk(data [][]rune, facing direction) [][2]int {
	src := findRune(data, 'S')
	dest := findRune(data, 'E')
	visited := map[[2]int]bool{}
	costs := map[[2]int]int{src:0}
	queue := map[[2]int]int{src:0}
	predecessors := map[[2]int]*[2]int{src:nil}

	var next [2]int
	for {
		//while queue not empty 
		if len(queue) == 0 {
			break
		}

		//pop next highest priority node
		queue, next = pop(queue)

		if next == dest {
			break
		}

		//do not visit same node twice
		if visited[next] {
			continue
		}

		//mark node as visited
		visited[next] = true

		//set current direction
		if predecessors[next] != nil {
			facing = getDirection(*predecessors[next], next)
		}

		//calculate costs of neighbouring nodes
		for _, n := range getNeighbours(data, next) {
			if data[n[1]][n[0]] == '#' {
				continue
			}

			currCost := costs[next] + 1
			if getDirection(next, n) != facing {
				currCost += 1000
			}

			//if we don't have cost for this node 
			//or it's less than what we have, add to queue
			if prev, ok := costs[n]; !ok || currCost < prev {
				costs[n] = currCost
				//not the best...
				queue[n] = costs[n] + heuristic(n, dest)
				pred := [2]int{next[0], next[1]}
				predecessors[n] = &pred
			}
		}
	}

	return makePath(predecessors, dest)
}

func makePath(predecessors map[[2]int]*[2]int, dest [2]int) [][2]int {
	path := [][2]int{dest}
	pred := predecessors[dest]
	for pred != nil {
		path = append(path, *pred)
		pred = predecessors[*pred]
	}
	slices.Reverse(path)
	return path
}


func heuristic(node, dest [2]int) int {
	a := math.Abs(float64(node[0]-dest[0]))
	b := math.Abs(float64(node[1]-dest[1])) 
	return int(a + b) * int(a + b)
}

func getDirection(curr, next [2]int) direction {
	if curr[1]-1 == next[1] {
		return NORTH 
	}
	if curr[1]+1 == next[1] {
		return SOUTH
	}
	if curr[0]+1 == next[0] {
		return EAST
	}
	if curr[0]-1 == next[0] {
		return WEST
	}
	panic("Fix your skill issues, Blizzard!")
}

var dirs = [4][2]int{{0,-1},{0,1},{1,0},{-1, 0}}

func getNeighbours(data [][]rune, node [2]int) [][2]int {
	neighbours := [][2]int{}
	for _, dir := range dirs {
		x, y := node[0]+dir[0], node[1]+dir[1]
		if y < 0 || y >= len(data) {
			continue
		}
		if x < 0 || x >= len(data[y]) {
			continue
		}
		neighbours = append(neighbours, [2]int{x, y})
	}
	return neighbours
}

func pop(queue map[[2]int]int) (map[[2]int]int, [2]int) {
	node, prio := [2]int{-1, -1}, math.MaxInt
	for k, v := range queue {
		if v < prio {
			node = k
			prio = v
		}
	}
	delete(queue, node)
	return queue, node
}

func sumPath(path [][2]int, facing direction) int {
	var sum int
	for i := 0; i < len(path)-1; i++ {
		newFacing := getDirection(path[i], path[i+1])
		sum += 1
		if newFacing != facing {
			sum += 1000
			facing = newFacing
		}
	}
	return sum
}

func findRune(data [][]rune, r rune) [2]int {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == r {
				return [2]int{j, i}
			}
		}
	}
	panic("Could not find rune position")
}
