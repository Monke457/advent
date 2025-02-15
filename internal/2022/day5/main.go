package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	//"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2022/day5.txt")

	crates, idx := parseCrates(data)

	for _, line := range data[idx:] {
		count, from, to := parseMove(line)
		move(&crates, from-1, to-1, count)
	}

	fmt.Println("Final top boxes:", getTopBoxes(crates))
}

func move(crates *[][]rune, from, to, count int) {
	stack := make([]rune, count)
	copy(stack, (*crates)[from][0:count])
	//slices.Reverse(stack)
	(*crates)[from] = (*crates)[from][count:]
	(*crates)[to] = append(stack, (*crates)[to]...)
}

func parseMove(moveStr string) (int, int, int) {
	parts := strings.Split(moveStr, " ")
	sCount := parts[1]
	sFrom := parts[3]
	sTo := parts[5]

	count, _ := strconv.Atoi(sCount)
	from, _ := strconv.Atoi(sFrom)
	to, _ := strconv.Atoi(sTo)

	return count, from, to
}

func parseCrates(data []string) ([][]rune, int) {
	crates := [][]rune{}
	for i, line := range data {
		if !strings.Contains(line, "[") {
			return crates, i+2
		}

		for i := 0; i < len(line); i++ {
			r := line[i]
			if !(r >= 'A' && r <= 'Z') {
				continue
			}
			idx := max(0, i/4) 
			for idx >= len(crates) {
				crates = append(crates, []rune{})
			}
			crates[idx] = append(crates[idx], rune(r))
		}
	}
	panic("data is bad and you are bad too")
}

func getTopBoxes(crates [][]rune) string{
	top := []rune{}
	for _, row := range crates {
		if len(row) == 0 {
			continue
		}
		top = append(top, row[0])
	}
	return string(top)
}
