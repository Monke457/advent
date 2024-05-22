package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

type blocks struct {
	cells   [][]int
	history [3]int8
}

const (
	north int8 = 1
	east       = 2
	south      = 3
	west       = 4
)

func main() {
	data := reader.FileTo2DIntArray("data/day17.txt")
	blocks := NewBlocks(data)
	blocks.Print()
	blocks.FindRoute()
}

func NewBlocks(cells [][]int) blocks {
	return blocks{cells: cells, history: [3]int8{}}
}

func (b blocks) Print() {
	for _, row := range b.cells {
		fmt.Printf("%d\n", row)
	}
}

func (b blocks) FindRoute() int {
	res := b.walk(0, 0, 0, math.MaxInt64)
	return res
}

func (b blocks) walk(row, col, h, m int) int {
	fmt.Printf("history %d\n", b.history)
	if row == len(b.cells)-1 && col == len(b.cells[0])-1 {
		fmt.Println("reached end at ", row, col, "min", min(h, m))
		return min(h, m)
	}
	if row < 0 || col < 0 || row >= len(b.cells) || col >= len(b.cells[0]) {
		fmt.Println("oob")
		return m
	}
	if row > 0 || col > 0 {
		fmt.Printf("adding %d to %d at [%d, %d]\n", b.cells[row][col], h, row, col)
		h += b.cells[row][col]
	}
	if h > m {
		return m
	}

	for _, i := range []int8{2, 3, 1, 4} {
		fmt.Println("loop", i)
		if b.history[0] == i && b.history[1] == i && b.history[2] == i {
			continue
		}
		switch i {
		case 2:
			if b.history[0] == 4 {
				continue
			}
			b.updateHistory(2)
			m = b.walk(row, col+1, h, m)
		case 3:
			if b.history[0] == 1 {
				continue
			}
			b.updateHistory(3)
			m = b.walk(row+1, col, h, m)
		case 1:
			if b.history[0] == 3 {
				continue
			}
			b.updateHistory(1)
			m = b.walk(row-1, col, h, m)
		case 4:
			if b.history[0] == 2 {
				continue
			}
			b.updateHistory(4)
			m = b.walk(row, col-1, h, m)
		}
	}

	fmt.Println("end of function")
	return m
}

func (b *blocks) updateHistory(val int8) {
	b.history[2] = b.history[1]
	b.history[1] = b.history[0]
	b.history[0] = val
}
