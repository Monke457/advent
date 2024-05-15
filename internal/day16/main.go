package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"sync"
)

type maze struct {
	cells [][]rune
}

type infos struct {
	m   map[[2]int][]int8
	mut sync.Mutex
}

var info = infos{m: map[[2]int][]int8{}}

func main() {
	data := reader.FileTo2DArray("data/day16.txt")
	maze := maze{cells: data}

	for _, row := range maze.cells {
		fmt.Printf("%c\n", row)
	}

	fmt.Println(solveFirstProblem(maze))
}

func solveFirstProblem(m maze) int {
	m.Walk(0, 0, 2)

	info.mut.Lock()
	for key, val := range info.m {
		fmt.Printf("cell %d: cardinals %d\n", key, val)
	}
	info.mut.Unlock()
	return 0
}

func (m maze) Walk(row, col, cardinal int) {
	for {
		switch cardinal {
		case 1:
			row--
			if row < 0 {
				return
			}
			if info.Update(row, col, 3) {
				return
			}
			// going up
			switch m.cells[row][col] {
			case '/':
				cardinal = 2
			case '\\':
				cardinal = 4
			case '-':
				m.Walk(row, col, 2)
				m.Walk(row, col, 4)
				return
			}
		case 2:
			col++
			if col == len(m.cells[0]) {
				return
			}
			if info.Update(row, col, 4) {
				return
			}
			// going right
			switch m.cells[row][col] {
			case '/':
				cardinal = 1
			case '\\':
				cardinal = 3
			case '|':
				m.Walk(row, col, 1)
				m.Walk(row, col, 3)
				return
			}
		case 3:
			row++
			if row == len(m.cells) {
				return
			}
			if info.Update(row, col, 1) {
				return
			}
			// going down
			switch m.cells[row][col] {
			case '/':
				cardinal = 4
			case '\\':
				cardinal = 2
			case '-':
				m.Walk(row, col, 2)
				m.Walk(row, col, 4)
				return
			}
		case 4:
			col--
			if col < 0 {
				return
			}
			if info.Update(row, col, 2) {
				return
			}
			// going left
			switch m.cells[row][col] {
			case '/':
				cardinal = 3
			case '\\':
				cardinal = 1
			case '|':
				m.Walk(row, col, 1)
				m.Walk(row, col, 3)
				return
			}
		}
	}
}

func (i *infos) Update(row, col int, cardinal int8) bool {
	i.mut.Lock()
	fmt.Println("updating info", i.m, "at", row, col, "with", cardinal)
	for _, val := range i.m[[2]int{row, col}] {
		if val == cardinal {
			return true
		}
	}
	i.m[[2]int{row, col}] = append(i.m[[2]int{row, col}], cardinal)
	fmt.Println("updated", i.m)
	i.mut.Unlock()
	return false
}
