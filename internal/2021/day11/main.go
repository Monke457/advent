package main

import (
	"advent/internal/pkg/grid"
	"advent/internal/pkg/reader"
	"fmt"

	"github.com/Monke457/printge"
)

func main() {
	octos := reader.FileTo2DIntArray("data/2021/day11.txt")

	i := 1
	count := 0
	flashes := 0

	for { 
		queue := map[[2]int]int{}
		for y := range octos {
			for x := range octos[y] {
				queue[[2]int{y, x}]++
			}
		}
		octos, flashes = step(octos, queue, 0)
		if allFlash(octos) {
			break
		}
		if i <= 100 {
			count += flashes
		}
		i++
	}

	fmt.Println("Flashes after 100 steps:", count)
	fmt.Println("first sync up:", i)
}

func allFlash(octos [][]int) bool {
	for _, row := range octos {
		for _, val := range row {
			if val != 0 {
				return false
			}
		}
	}
	return true
}

func step(octos [][]int, queue map[[2]int]int, flashes int) ([][]int, int) {
	if len(queue) == 0 {
		return octos, flashes
	}
	flashed := map[[2]int]bool{}
	next := map[[2]int]int{}
	for node, val := range queue {
		y, x := node[0], node[1]
		octos[y][x] = octos[y][x] + val
		if octos[y][x] > 9 {
			octos[y][x] = 0
			flashed[node] = true
		}
	}
	for node := range flashed {
		flashes++
		for _, n := range grid.GetNeighbours(octos, node) {
			if octos[n[0]][n[1]] == 0 {
				continue
			}
			next[n]++
		}
	}
	return step(octos, next, flashes)
}

func printoctos(octos [][]int, step int) {
	fmt.Println("flash", step)
	for y := range octos {
		for x := range octos[y] {
			if octos[y][x] == 0 {
				printge.Print(fmt.Sprintf("%d", octos[y][x]), printge.CYAN)
			} else {
				fmt.Printf("%d", octos[y][x])
			}
		}
		fmt.Println()
	}
}
