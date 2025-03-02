package main

import (
	"advent/internal/pkg/grid"
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"

	"github.com/Monke457/printge"
)

func main() {
	data := reader.FileTo2DIntArray("data/2021/day15.txt")

	start := [2]int{0,0}
	end := [2]int{len(data)-1,len(data[0])-1}

	path := findBestPath(data, start, end)
	risk := calculateRisk(data, path)
	//printPath(data, path)
	fmt.Println("\nRisk of best path:", risk)

	// caveman solution ugg ugg make fire with laptop
	expanded := expandData(data)
	end = [2]int{len(expanded)-1,len(expanded[0])-1}
	path = findBestPath(expanded, start, end)
	risk = calculateRisk(expanded, path)
	fmt.Println("\nRisk of best path expanded:", risk)
}

func expandData(data [][]int) [][]int {
	height := len(data)
	full := [][]int{}
	for i := range 5 {
		for j := range 5 {
			temp := createIncrement(data, i+j)
			for idx, tempRow := range temp {
				row := i * height + idx
				for row >= len(full) {
					full = append(full, []int{})
				}
				full[row] = append(full[row], tempRow...)
			}
		}
	}
	return full
}

func createIncrement(data [][]int, i int) [][]int {
	inc := make([][]int, len(data))
	for y := range data {
		for _, val := range data[y] {
			val = val+i
			if val > 9 {
				val = val % 10 + 1
			}
			inc[y] = append(inc[y], val)
		}
	}
	return inc
}

func findBestPath(data [][]int, start, end [2]int) [][2]int {
	queue := map[[2]int]int{start:0}
	costs := map[[2]int]int{start:0}
	predecessors := map[[2]int]*[2]int{start:nil}
	
	var curr [2]int
	for len(queue) > 0 {
		curr = pop(&queue)
		if curr == end {
			break
		}
		for _, n := range grid.GetNeighboursCont(data, curr) {
			cost := costs[curr] + data[n[0]][n[1]] 

			if cached, ok := costs[n]; !ok || cost < cached {
				costs[n] = cost
				queue[n] = cost + heuristic(n, end)
				temp := [2]int{curr[0], curr[1]}
				predecessors[n] = &temp
			}
		}
	}

	path := makePath(predecessors, end)
	return path
}

func printPath(data [][]int, path [][2]int) {
	printge.Print("Shortest path\n\n", printge.WHITE)
	for y := range data {
		for x, val := range data[y] {
			if slices.Contains(path, [2]int{y, x}) {
				printge.Print(fmt.Sprintf("%d", val), printge.WHITE)
			} else {
				fmt.Printf("%d", val)
			}
		}
		fmt.Println()
	}
}

func heuristic(a, b [2]int) int {
	return int(math.Abs(float64(a[0] - b[0])) + math.Abs(float64(a[1] - b[1])))
}


func pop(queue *map[[2]int]int) [2]int {
	var node [2]int
	val := math.MaxInt
	for k, v := range *queue {
		if v < val {
			val = v
			node = k
		}
	}
	delete(*queue, node)
	return node
}

func makePath(predecessors map[[2]int]*[2]int, dest [2]int) [][2]int {
	path := [][2]int{dest}
	node := predecessors[dest]
	for node != nil {
		path = append(path, *node)
		node = predecessors[*node]
	}
	slices.Reverse(path)
	return path
}

func initQueue(data [][]int) map[[2]int]int {
	queue := map[[2]int]int{}
	for y := range data {
		for x := range data[y] {
			queue[[2]int{y, x}] = data[y][x] 
		}
	}
	return queue
}

func calculateRisk(data [][]int, path [][2]int) int {
	risk := 0
	for _, cell := range path[1:] {
		risk += data[cell[0]][cell[1]]
	}
	return risk
}
