package main

import (
	"advent/internal/pkg/reader"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type move struct {
	pipe  byte
	pos   [2]int
	steps int
	seed  int
}

var north [3]byte = [3]byte{
	124, 76, 74,
}
var east [3]byte = [3]byte{
	45, 70, 76,
}
var south [3]byte = [3]byte{
	124, 55, 70,
}
var west [3]byte = [3]byte{
	45, 55, 74,
}

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	lines := reader.FileToArray("data/day10.txt")
	start := findStart(lines)

	moves, result := solveFirstProblem(start, lines)
	fmt.Println(result)
	fmt.Println(solveSecondProblem(moves, lines))
}

func solveFirstProblem(start [2]int, data []string) ([]move, int) {
	pos := getStartMoves(data, start)

	c := make(chan move)
	for i, p := range pos {
		m := []move{
			{pipe: data[start[0]][start[1]], pos: start, seed: i},
		}
		go run(data, m, p, i, c)
	}

	moves := []move{
		{pipe: data[start[0]][start[1]], pos: start},
		{pipe: data[pos[0][0]][pos[0][1]], pos: pos[0], steps: 1, seed: 0},
		{pipe: data[pos[1][0]][pos[1][1]], pos: pos[1], steps: 1, seed: 1},
	}

	for {
		val := <-c
		for _, m := range moves {
			if m.pos == val.pos && m.steps == val.steps {
				return moves, m.steps
			}
		}
		moves = append(moves, val)
	}
}

func solveSecondProblem(moves []move, data []string) int {
	slices.SortFunc(moves, func(a, b move) int {
		if a.seed == b.seed {
			if a.seed == 0 {
				return cmp.Compare(a.steps, b.steps)
			}
			return cmp.Compare(b.steps, a.steps)
		}
		if a.seed > b.seed {
			return 1
		} else {
			return -1
		}
	})
	for _, m := range moves {
		fmt.Printf("\npipe: %c, pos: [%d,%d], steps: %d, seed. %d", m.pipe, m.pos[0], m.pos[1], m.steps, m.seed)
	}
	fmt.Println()
	return len(data)*len(data[0]) - len(moves)
}

func run(data []string, moves []move, p [2]int, seed int, c chan<- move) {
	steps := 1
	moves = append(moves, move{pipe: data[p[0]][p[1]], pos: p, steps: steps, seed: seed})
loop:
	for {
		for i, d := range directions {
			row := moves[steps].pos[0] + d[0]
			col := moves[steps].pos[1] + d[1]

			if row < 0 || col < 0 || row >= len(data) || col >= len(data[row]) {
				continue
			}

			next := data[row][col]
			if next == 46 || contains(moves, row, col) {
				continue
			}

			curr := moves[steps].pipe
			if m := getMove(i, row, col, next, curr); m != nil {
				steps++
				m.steps = steps
				m.seed = seed
				moves = append(moves, *m)
				c <- *m
				continue loop
			}
		}
		return
	}
}

func getMove(i, row, col int, next, curr byte) *move {
	switch i {
	case 0:
		if valid(next, south[:]...) && valid(curr, north[:]...) {
			return &move{pipe: next, pos: [2]int{row, col}}
		}
		return nil
	case 1:
		if valid(next, west[:]...) && valid(curr, east[:]...) {
			return &move{pipe: next, pos: [2]int{row, col}}
		}
		return nil
	case 2:
		if valid(next, north[:]...) && valid(curr, south[:]...) {
			return &move{pipe: next, pos: [2]int{row, col}}
		}
		return nil
	case 3:
		if valid(next, east[:]...) && valid(curr, west[:]...) {
			return &move{pipe: next, pos: [2]int{row, col}}
		}
		return nil
	}
	return nil
}

func contains(moves []move, row, col int) bool {
	for _, m := range moves {
		if m.pos[0] == row && m.pos[1] == col {
			return true
		}
	}
	return false
}

func valid(r byte, vals ...byte) bool {
	for _, b := range vals {
		if r == b {
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
		next := data[row][col]
		if next == 46 {
			continue
		}
		switch i {
		case 0:
			if valid(next, south[:]...) {
				result = append(result, [2]int{row, col})
			}
		case 1:
			if valid(next, west[:]...) {
				result = append(result, [2]int{row, col})
			}
		case 2:
			if valid(next, north[:]...) {
				result = append(result, [2]int{row, col})
			}
		case 3:
			if valid(next, east[:]...) {
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
