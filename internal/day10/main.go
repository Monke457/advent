package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
	"time"
)

type move struct {
	pos   [2]int
	steps int
}

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

func main() {
	lines := reader.FileToArray("data/day10.txt")
	start := findStart(lines)

	fmt.Println(solveFirstProblem(start, lines))
}

func solveFirstProblem(start [2]int, data []string) int {
	pos := getStartMoves(data, start)

	c := make(chan move)
	for _, p := range pos {
		go run(data, start, p, c)
	}

	moves := []move{
		{pos: start},
		{pos: pos[0], steps: 1},
		{pos: pos[1], steps: 1},
	}
	for {
		select {
		case val := <-c:
			for _, m := range moves {
				if m.pos == val.pos && m.steps == val.steps {
					return m.steps
				}
			}
			moves = append(moves, val)
		case <-time.After(5 * time.Second):
			for _, m := range moves {
				if m.steps == 10 {
					fmt.Println(m)
				}
			}
			return 0
		}
	}
}

func run(data []string, start, pos [2]int, c chan<- move) {
	steps := 1
	path := [][2]int{start, pos}
loop:
	for {
		steps++
		for i, d := range directions {
			row := pos[0] + d[0]
			col := pos[1] + d[1]
			r := data[row][col]
			if r == 46 {
				continue
			}
			switch i {
			case 0:
				if r == 124 || r == 55 || r == 70 {
					if contains(path, row, col) {
						continue
					}
					pos[0] = row
					pos[1] = col
					path = append(path, pos)
					c <- move{pos: [2]int{row, col}, steps: steps}
					continue loop
				}
			case 1:
				if r == 45 || r == 55 || r == 74 {
					if contains(path, row, col) {
						continue
					}
					pos[0] = row
					pos[1] = col
					path = append(path, pos)
					c <- move{pos: [2]int{row, col}, steps: steps}
					continue loop
				}
			case 2:
				if r == 45 || r == 70 || r == 76 {
					if contains(path, row, col) {
						continue
					}
					pos[0] = row
					pos[1] = col
					path = append(path, pos)
					c <- move{pos: [2]int{row, col}, steps: steps}
					continue loop
				}
			case 3:
				if r == 124 || r == 76 || r == 74 {
					if contains(path, row, col) {
						continue
					}
					pos[0] = row
					pos[1] = col
					path = append(path, pos)
					c <- move{pos: [2]int{row, col}, steps: steps}
					continue loop
				}
			}
		}
	}
}

func contains(arr [][2]int, row, col int) bool {
	for _, v := range arr {
		if v[0] == row && v[1] == col {
			return true
		}
	}
	return false
}

func getStartMoves(data []string, pos [2]int) [][2]int {
	result := [][2]int{}
	for i, d := range directions {
		row := pos[0] + d[0]
		col := pos[1] + d[1]
		r := data[row][col]
		if r == 46 {
			continue
		}
		switch i {
		case 0:
			if r == 124 || r == 55 || r == 70 {
				result = append(result, [2]int{row, col})
			}
		case 1:
			if r == 45 || r == 55 || r == 74 {
				result = append(result, [2]int{row, col})
			}
		case 2:
			if r == 45 || r == 70 || r == 76 {
				result = append(result, [2]int{row, col})
			}
		case 3:
			if r == 124 || r == 76 || r == 74 {
				result = append(result, [2]int{row, col})
			}
		}
	}
	return result
}

func findStart(lines []string) [2]int {
	var start [2]int
	for row, l := range lines {
		if strings.Contains(l, "S") {
			for col, r := range l {
				if r == 83 {
					start = [2]int{row, col}
					break
				}
			}
			break
		}
	}
	return start
}

func equal(pos ...[2]int) bool {
	for i := 0; i < len(pos)-1; i++ {
		if pos[i] != pos[i+1] {
			return false
		}
	}
	return true
}
