package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

const(
	draw int8 = 0
	s_row = 1
	s_col = 2
)

type screen struct {
	cells [6][50]bool
}

type command struct {
	typ int8
	val [2]int
}

func main() {
	data := reader.FileToArray("data/2016/day8.txt")

	screen := screen{ cells: [6][50]bool{} }
	for _, line := range data {
		cmd := parseCommand(line)
		switch cmd.typ {
		case 0:
			screen = screen.draw(cmd.val)
		case 1:
			screen = screen.ShiftRow(cmd.val)
		case 2:
			screen = screen.ShiftCol(cmd.val)
		}
	}

	screen.print()
	fmt.Println(screen.countLit())
}

func (s screen) countLit() int {
	result := 0
	for _, row := range s.cells {
		for _, col := range row {
			if col {
				result++
			}
		}
	}
	return result
}

func (s screen) draw(val [2]int) screen {
	for i := 0; i < val[0]; i++ {
		for j := 0; j < val[1]; j++ {
			s.cells[j][i] = true
		}
	}
	return s
}

func (s screen) ShiftRow(val [2]int) screen {
	row, pos := val[0], len(s.cells[0]) - val[1]
	start, end := s.cells[row][:pos], s.cells[row][pos:]	
	newRow := append(end, start[:]...)
	for i := range newRow {
		s.cells[row][i] = newRow[i] 
	}
	return s
}

func (s screen) ShiftCol(val [2]int) screen {
	l := len(s.cells)
	newCol := make([]bool, l)
	for i := 0; i < l; i++ {
		idx := (i + val[1]) % l
		newCol[idx] = s.cells[i][val[0]]
	}
	for i := 0; i < l; i++ {
		s.cells[i][val[0]] = newCol[i]
	}
	return s
}

func parseCommand(str string) command {
	cmd := command{}
	var v1, v2 string
	if strings.Contains(str, "rect") {
		cmd.typ = 0
		v1, v2, _ = strings.Cut(strings.Split(str, " ")[1], "x")

	} else {
		if strings.Contains(str, "row") {
			cmd.typ = 1
		} else {
			cmd.typ = 2
		}
		v1, v2, _ = strings.Cut(str, " by ")
		v1 = strings.Split(v1, "=")[1]
	}

	v1i, e1 := strconv.Atoi(v1)
	v2i, e2 := strconv.Atoi(v2)
	if e1 != nil || e2 != nil {
		panic("conversion error")
	}
	cmd.val = [2]int{v1i, v2i}

	return cmd
}

func (s screen) print() {
	for _, row := range s.cells {
		for _, cell := range row {
			if cell {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
