package main

import (
	"advent/pkg/reader"
	"fmt"
)

type maze struct {
	cells [][]rune
}

type infos struct {
	m map[[2]int][]int8
}

func main() {
	data := reader.FileTo2DArray("data/2023/day16.txt")
	maze := maze{cells: data}

	fmt.Println("first:", solveFirstProblem(maze))
	fmt.Println("second:", solveSecondProblem(maze))
}

func solveFirstProblem(m maze) int {
	c := make(chan int)
	go func() {
		info := infos{m: map[[2]int][]int8{}}
		m.Walk(info, 0, -1, 2)
		c <- len(info.m)
	}()

	return <-c
}

func solveSecondProblem(m maze) int {
	c := make(chan int)

	count := 0
	for col := range m.cells[0] {
		go m.StartWalk(-1, col, 3, c)
		go m.StartWalk(len(m.cells), col, 1, c)
		count += 2
	}

	for row := range m.cells {
		go m.StartWalk(row, -1, 2, c)
		go m.StartWalk(row, len(m.cells[0]), 4, c)
		count += 2
	}

	res := 0
	for range count {
		tiles := <-c
		if tiles > res {
			res = tiles
		}
	}
	return res
}

func (m maze) StartWalk(row, col, dir int, c chan int) {
	info := infos{m: map[[2]int][]int8{}}
	m.Walk(info, row, col, dir)
	c <- len(info.m)
	return
}

func (m maze) Walk(info infos, row, col, cardinal int) {
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
				m.Walk(info, row, col, 2)
				m.Walk(info, row, col, 4)
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
				m.Walk(info, row, col, 1)
				m.Walk(info, row, col, 3)
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
				m.Walk(info, row, col, 2)
				m.Walk(info, row, col, 4)
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
				m.Walk(info, row, col, 1)
				m.Walk(info, row, col, 3)
				return
			}
		}
	}
}

func (i *infos) Update(row, col int, cardinal int8) bool {
	for _, val := range i.m[[2]int{row, col}] {
		if val == cardinal {
			return true
		}
	}
	i.m[[2]int{row, col}] = append(i.m[[2]int{row, col}], cardinal)
	return false
}
