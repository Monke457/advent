package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

type blocks struct {
	cells [][]int
}

var directions = [4]int{1, 2, 3, 4}

func main() {
	data := reader.FileTo2DIntArray("data/day17.txt")
	blocks := blocks{cells: data}
	for _, row := range blocks.cells {
		fmt.Printf("%d\n", row)
	}
	blocks.StartWalk()
}

func (b blocks) StartWalk() {
	c := make(chan []int)
	go func() {
		c <- b.Walk(0, 1, 0, [3]int{2}, map[[2]int]bool{}, []int{})
	}()
	/*
		go func() {
			c <- b.Walk(1, 0, 0, [3]int{4}, []int{})
		}()
	*/
	results := <-c
	//	results = append(results, (<-c)[:]...)
	fmt.Printf("Heat values - %d\n", results)
}

func (b blocks) Walk(row, col, heat int, dirs [3]int, seen map[[2]int]bool, results []int) []int {
	if row == len(b.cells)-1 && col == len(b.cells[0])-1 {
		results = append(results, heat)
		fmt.Println("adding heat value to array", heat)
		return results
	}
	if row < 0 || col < 0 || row == len(b.cells) || col == len(b.cells[0]) {
		fmt.Println("oob", row, col)
		return results
	}
	if seen[[2]int{row, col}] {
		fmt.Println("seen")
		return results
	}

	heat += b.cells[row][col]
	seen[[2]int{row, col}] = true

	for _, d := range directions {
		if d == dirs[0] && d == dirs[1] && d == dirs[2] {
			continue
		}
		switch d {
		case 1:
			if dirs[0] == 3 {
				continue
			}
			dirs := shiftArray(dirs, d)
			res := b.Walk(row-1, col, heat, dirs, seen, results)
			results = append(results, res[:]...)
		case 2:
			if dirs[0] == 4 {
				continue
			}
			dirs := shiftArray(dirs, d)
			res := b.Walk(row, col+1, heat, dirs, seen, results)
			results = append(results, res[:]...)
		case 3:
			if dirs[0] == 1 {
				continue
			}
			dirs := shiftArray(dirs, d)
			res := b.Walk(row+1, col, heat, dirs, seen, results)
			results = append(results, res[:]...)
		case 4:
			if dirs[0] == 2 {
				continue
			}
			dirs := shiftArray(dirs, d)
			res := b.Walk(row, col-1, heat, dirs, seen, results)
			results = append(results, res[:]...)
		}
	}
	return results
}

func shiftArray(arr [3]int, val int) [3]int {
	arr[2] = arr[1]
	arr[1] = arr[0]
	arr[0] = val
	return arr
}
