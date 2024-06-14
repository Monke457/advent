package main

import (
	"fmt"
	"go/types"
	"math"
)

func main() {
	input := 312051
	solveFirst(input)
	solveSecond(input)
}

func solveFirst(num int) {
	rt, sq := findSquare(num)
	result := findDistance(num, rt, sq)
	fmt.Println("First:", result)
}

func solveSecond(num int) {
	var err *types.Error
	start := 1
	grid := [][]*int{ { &start } }
	row, col := 0, 0 
	val := 1
	for val < num {
		row, col, err = move(grid, row, col)
		if err != nil {
			if err.Msg == "NR" {
				grid = updateGrid(grid, row-1, col-1)
				row += 1
			} else {
				panic(err)
			}
		}

		grid = updateGrid(grid, row, col)
		row = max(row, 0)
		col = max(col, 0)

		val = 0	
		for _, neighbour := range getNeighbours(grid, row, col) {
			val += neighbour
		}
		n := val
		grid[row][col] = &n
	}
	printGrid(grid)
	fmt.Println("Second:", val)
}

var neighbours = [][2]int{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func getNeighbours(grid [][]*int, row, col int) []int {
	res := []int{}
	for _, n := range neighbours {
		nr, nc := row + n[0], col + n[1]
		if nr >= len(grid) || nr < 0 || nc >= len(grid[nr]) || nc < 0 {
			continue
		}
		val := grid[row + n[0]][col + n[1]]
		if val == nil {
			continue
		}
		res = append(res, *val)
	}
	return res 
}

func move(grid [][]*int, row, col int) (int, int, *types.Error) {
	h, l := len(grid), getLen(grid)

	if row + 1 == h && col + 1 == l {
		if colContainsNil(grid, col) {
			return row - 1, col, nil
		}
		if h == 1 {
			return row, l, &types.Error { Msg: "NR" }
		}
		return row, l, nil 
	}

	if row == 0 && col == 0 {
		if colContainsNil(grid, col) {
			return row + 1, col, nil
		}
		return row, col - 1, nil 
	}

	if row == 0 {
		if rowContainsNil(grid, row) {
			return row, col - 1, nil
		}
		return row - 1, col, nil
	}

	if row + 1 == h {
		if rowContainsNil(grid, row) {
			return row, col + 1, nil
		}
		return row + 1, col, nil
	}

	if col + 1 == l {
		return row - 1, col, nil
	}

	if col == 0 {
		return row + 1, col, nil
	}

	return row, col, &types.Error{ Msg: "Do not know where to go" }
}

func colContainsNil(grid [][]*int, col int) bool {
	for _, row := range grid {
		if row[col] == nil {
			return true
		}
	}
	return false
}

func rowContainsNil(grid [][]*int, row int) bool {
	for _, val := range grid[row] {
		if val == nil {
			return true
		}
	}
	return false
}

func updateGrid(grid [][]*int, row, col int) [][]*int {
	h, l := len(grid), getLen(grid)

	if row > h || col > l || (row == h && col == l) {
		panic("something has gone wrong with the movement")
	}

	if row == h || row < 0 {
		newRow := []*int{}
		for i := 0; i < l; i++ {
			newRow = append(newRow, nil)
		}
		if row == h {
			grid = append(grid, newRow) 
		} else {
			grid = append([][]*int{ newRow }, grid[:]...) 
		}
	}

	if col == l || col < 0 {
		for i := range grid {
			if col == l {
				grid[i] = append(grid[i], nil)
			} else {
				grid[i] = append([]*int{nil}, grid[i][:]...)
			}
		}
	}

	return grid
}

func findDistance(n, rt, sq int) int {
	l := max(rt-1, 1)
	c := (rt-2) * (rt-2) + 1
	low, high := sq, 1
	for low > n {
		low = max(low-l, c) 
	}
	for high < n {
		high += l 
	}
	mid := high - (l >> 1)
	diff := int(math.Abs(float64(mid - n)))
	return (l >> 1) + diff 
}

func findSquare(n int) (int, int) {
	res := 1
	sq := 0
	for sq < n {
		sq = res * res 
		res += 2
	}
	res -= 2
	return res, sq
}

func printGrid(grid [][]*int) {
	l, h := getLen(grid), len(grid)
	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			c := grid[i][j]
			if c == nil {
				fmt.Printf("\t")
			} else {
				fmt.Printf("%d\t", *c)
			}
		}
		fmt.Println()
	}
}

func getLen(grid [][]*int) int {
	res := 0
	for _, row := range grid {
		if len(row) > res {
			res = len(row)
		}
	}
	return res
}
