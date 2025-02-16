package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileTo2DIntArray("data/2022/day8.txt")

	visible := map[[2]int]bool{}
	for y := range data {
		loop:
		for x := range data[y] {
			for _, dir := range dirs {
				if isVisible(data, dir, y, x, data[y][x]) {
					visible[[2]int{y, x}] = true
					continue loop
				}
			}
		}
	}
	print(data, visible)
	fmt.Println("visible trees:", len(visible))


	scenicScores := map[[2]int]int{}

	for y := range data {
		for x := range data[y] {
			scenicScores[[2]int{y,x}] = calculateScore(data, y, x)
		}
	}

	best := 0
	for _, score := range scenicScores {
		if score > best {
			best = score
		}
	}
	fmt.Println("Best scenic score:", best)
}

func calculateScore(trees [][]int, y, x int) int {
	total := 1
	if y == 0 || y == len(trees)-1 {
		return 0
	}
	if x == 0 || x == len(trees[y])-1 {
		return 0
	}

	for _, dir := range dirs {
		count := 1
		cY := y + dir[0]
		cX := x + dir[1]

		for cY > 0 && cX > 0 && cY < len(trees)-1 && cX < len(trees[cY])-1 && trees[cY][cX] < trees[y][x] {
			cY += dir[0]
			cX += dir[1]
			count++
		}

		total *= count
	}
	return total
}

func print(trees [][]int, visible map[[2]int]bool) {
	for y := range trees {
		for x := range trees[y] {
			if visible[[2]int{y, x}] {
				fmt.Printf("\033[0;32m%d\033[0m", trees[y][x])
			} else {
				fmt.Printf("\033[0;31m%d\033[0m", trees[y][x])
			}
		}
		fmt.Println()
	}
}

func isVisible(trees [][]int, dir [2]int, y, x, val int) bool {
	if y == 0 || y == len(trees)-1 {
		return true
	}
	if x == 0 || x == len(trees[y])-1 {
		return true
	}
	nY, nX := y + dir[0], x + dir[1]
	if trees[nY][nX] >= val {
		return false
	}
	return isVisible(trees, dir, nY, nX, val)
}

var dirs = [][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
