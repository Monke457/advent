package main

import (
	"advent/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2016/day6.txt")

	letters := make([]map[rune]int, len(data[0]))
	for _, row := range data {
		for j, col := range row {
			if letters[j] == nil {
				letters[j] = make(map[rune]int)
			}
			letters[j][col]++
		}
	}
	solveFirstProblem(letters)
	solveSecondProblem(letters)
}

func solveFirstProblem(letters []map[rune]int) {
	result := strings.Builder{}
	for _, m := range letters {
		count := 0
		char := ' '
		for k, v := range m {
			if v > count {
				char = k
				count = v
			}
		}
		result.WriteRune(char)
	}
	fmt.Println(result.String())
}

func solveSecondProblem(letters []map[rune]int) {
	result := strings.Builder{}
	for _, m := range letters {
		count := 0
		char := ' '
		for k, v := range m {
			if count == 0 || v < count {
				char = k
				count = v
			}
		}
		result.WriteRune(char)
	}
	fmt.Println(result.String())
}
