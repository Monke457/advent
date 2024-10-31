package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2015/day12.txt")
	fmt.Println("First problem:", solveFirstProblem(data))
	fmt.Println("Second problem:", solveSecondProblem(data))
}

func solveFirstProblem(data string) int {
	return sumInts(data)
}

func solveSecondProblem(data string) int {
	return sumWithoutRed(data)
}

func sumWithoutRed(data string) int {
	return 0 
}

func sumInts(data string) int {
	var total int
	sb := strings.Builder{}
	for _, b := range data {
		if b == '-' {
			sb.WriteRune(b)
		} else if b > 47 && b < 58 {
			sb.WriteRune(b)
		} else if sb.Len() > 0 {
			val, err := strconv.Atoi(sb.String())
			if err != nil {
				panic(err)
			}
			total += val
			sb.Reset()
		}
	}
	return total
}
