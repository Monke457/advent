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

type state struct {
	pos [2]int
	facing direction
}

func main() {
	data := reader.FileTo2DArray("data/2024/day16.txt")

	paths := aStarSearch(data, math.MaxInt)
	best := sumPath(paths[0], EAST)
	reduced := reducePaths(paths)

	fmt.Println("Bet score possible:", best)
	fmt.Println("Tiles that are part of at least one winning path:", len(reduced))
}

func reducePaths(paths [][][2]int) map[[2]int]bool {
	res := map[[2]int]bool{}
	for _, path := range paths {
		for _, tile := range path {
			res[tile] = true
		}
	}
	return res
}

func aStarSearch(data [][]rune, best int) [][][2]int {
	src := state{findRune(data, 'S'), EAST}
	dest := findRune(data, 'E')
	queue := map[state]int{src:0}
	visited := map[state]bool{}
	destStates := []state{}
	predecessors := map[state]map[state]bool{src:{}}

	var next state
	var cost int

	for len(queue) > 0 {
		queue, next, cost = pop(queue)

		if cost > best {
			continue
		}

		if next.pos == dest {
			if cost < best {
				best = cost
				clear(predecessors)
				queue[src] = 0
			}
			destStates = append(destStates, next)
			clear(visited)
			continue
		}

		if visited[next] {
			continue
		}
		visited[next] = true

		for _, n := range getNeighbours(data, next) {
			if visited[n] {
				continue
			}

			currCost := cost + 1
			if n.facing != next.facing {
				currCost += 1000
			}
			if currCost > best {
				continue
			} 

			if prev, ok := queue[n]; !ok || currCost < prev {
				queue[n] = currCost
			}
			if _, ok := predecessors[n]; !ok {
				predecessors[n] = map[state]bool{}
			}
			predecessors[n][next] = true
		}
	}

	finalPaths := getFinalPaths(predecessors, destStates, src.pos, best)
	return finalPaths
}

func getFinalPaths(preds map[state]map[state]bool, targetStates []state, start [2]int, best int) [][][2]int {
	paths := [][][2]int{}
	for _, target := range targetStates {
		paths = append(paths, backtrack(preds, map[[2]int]bool{}, [][2]int{target.pos}, target, start, 0, best)...)
	}
	return paths 
}

func backtrack(preds map[state]map[state]bool, visited map[[2]int]bool, path [][2]int, curr state, start [2]int, cost, best int) [][][2]int {
	if curr.pos == start {
		return [][][2]int{path}
	}
	allPaths := [][][2]int{}
	visited[curr.pos] = true
	for prev := range preds[curr] {
		if visited[prev.pos] {
			continue
		}

		moveCost := 1
		if prev.facing != curr.facing {
			moveCost += 1000
		}
		if cost + moveCost > best {
			continue
		}

		path = append([][2]int{prev.pos}, path...)
		allPaths = append(allPaths, backtrack(preds, visited, path, prev, start, cost+moveCost, best)...)
		path = path[1:]
	}
	delete(visited, curr.pos)
	return allPaths
}

var dirs = map[direction][2]int{NORTH:{0,-1},SOUTH:{0,1},EAST:{1,0},WEST:{-1, 0}}

func getNeighbours(data [][]rune, s state) []state {
	neighbours := []state{}
	for facing, coords := range dirs {
		x, y := s.pos[0]+coords[0], s.pos[1]+coords[1]
		if y < 0 || y >= len(data) {
			continue
		}
		if x < 0 || x >= len(data[y]) {
			continue
		}
		if data[y][x] == '#' {
			continue
		}
		pos := [2]int{x, y}
		neighbours = append(neighbours, state{pos, facing})
	}
	return neighbours
}

func pop(queue map[state]int) (map[state]int, state, int) {
	s, prio := state{}, math.MaxInt
	for k, v := range queue {
		if v < prio {
			s = k
			prio = v
		}
	}
	delete(queue, s)
	return queue, s, prio
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
	panic(fmt.Errorf("Fix your skill issues, Blizzard! %d %d", curr, next))
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

func draw(data [][]rune, path [][2]int) {
	for y := range data {
		for x := range data[y] {
			if slices.Contains(path, [2]int{x, y}) {
				fmt.Printf("0")
			} else {
				fmt.Printf("%c", data[y][x])
			}
		}
		fmt.Println()
	}
}

