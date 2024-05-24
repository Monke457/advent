package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
)

var neighbours [][]int = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	fmt.Println(solveFirstProblem())
	fmt.Println(solveSecondProblem())
}

func solveFirstProblem() int {
	lines := reader.FileToArray("data/2023/day3.txt")

	parts := []int{}
	for i, l := range lines {
		for j := 0; j < len(l); j++ {
			if j >= len(l) {
				continue
			}
			if l[j] > 57 || l[j] < 48 {
				continue
			}

			var p string
			isPart := false
			for l[j] < 58 && l[j] > 47 {
				if !isPart && isSymbolAdjacent(lines, i, j) {
					isPart = true
				}
				p += fmt.Sprintf("%c", l[j])
				j++
				if j >= len(l) {
					break
				}
			}

			if !isPart {
				continue
			}

			part, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			parts = append(parts, part)
		}
	}

	sum := 0
	for _, p := range parts {
		sum += p
	}

	return sum
}

func solveSecondProblem() int {
	lines := reader.FileToArray("data/2023/day3.txt")

	gears := map[string][]int{}
	for i, l := range lines {
		for j := 0; j < len(l); j++ {
			if j >= len(l) {
				continue
			}
			if l[j] > 57 || l[j] < 48 {
				continue
			}

			var r int
			var c int
			var g string
			isGear := false
			for l[j] < 58 && l[j] > 47 {
				if !isGear {
					isGear, r, c = isGearAdjacent(lines, i, j)
				}
				g += fmt.Sprintf("%c", l[j])
				j++
				if j >= len(l) {
					break
				}
			}

			if !isGear {
				continue
			}

			gear, err := strconv.Atoi(g)
			if err != nil {
				panic(err)
			}
			key := fmt.Sprintf("r%dc%d", r, c)
			gears[key] = append(gears[key], gear)
		}
	}

	sum := 0
	for _, val := range gears {
		if len(val) != 2 {
			continue
		}
		sum += val[0] * val[1]
	}

	return sum
}

func isGearAdjacent(lines []string, row, col int) (bool, int, int) {
	for _, n := range neighbours {
		r := row + n[0]
		c := col + n[1]

		if isOOB(lines, r, c) {
			continue
		}

		if lines[r][c] == 42 {
			return true, r, c
		}
	}
	return false, 0, 0
}

func isSymbolAdjacent(lines []string, row, col int) bool {
	for _, n := range neighbours {
		r := row + n[0]
		c := col + n[1]

		if isOOB(lines, r, c) {
			continue
		}

		if isSymbol(lines[r][c]) {
			return true
		}
	}
	return false
}

func isSymbol(b byte) bool {
	if b == 46 {
		return false
	}
	if b < 58 && b > 47 {
		return false
	}
	return true
}

func isOOB(lines []string, row, col int) bool {
	if row < 0 || row >= len(lines) {
		return true
	}
	if col < 0 || col >= len(lines[row]) {
		return true
	}
	return false
}
