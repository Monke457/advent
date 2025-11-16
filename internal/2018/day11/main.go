package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

type powerlevel struct {
	coords [2]int
	size int
	sum int
}

func main() {
	sn := reader.FileToIntArray("data/2018/day11.txt")[0]

	w, h := 300, 300
	grid := initGrid(w, h, sn)

	var best powerlevel
	for i := range w {
		coord, sum := findBest(grid, w, h, i+1)
		if sum == 0 {
			break
		}
		if sum > best.sum {
			pl := powerlevel{coord, i+1, sum}
			best = pl
		}
		fmt.Printf("\rDone: %d", i+1)
	}

	fmt.Println("\nBest:", best)
}

func findBest(grid [][]int, w, h, size int) ([2]int, int) {
	var best int
	var coord [2]int
	for y := range w-size {
		for x := range h-size {
			sum := sumPower(grid, x, y, size)
			if sum > best {
				best = sum
				coord = [2]int{x+1,y+1}
			}
		}
	}
	return coord, best 
}

func sumPower(grid [][]int, x, y, size int) int {
	sum := 0
	for i := range size {
		for j := range size {
			sum += grid[y+i][x+j]
		}
	}
	return sum
}

func calculatePower(x, y, sn int) int {
	id := x + 10
	pow := id * y
	pow += sn
	pow *= id
	pow = (pow / 100) % 10
	pow -= 5
	return pow
}

func initGrid(w, h, sn int) [][]int { 
	grid := make([][]int, h)
	for y := range grid {
		row := make([]int, w)
		for x := range len(row) {
			row[x] = calculatePower(x+1,y+1,sn)
		}
		grid[y] = row
	}
	return grid
}
