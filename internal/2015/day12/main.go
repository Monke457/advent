package main

import (
	"advent/pkg/reader"
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
	var total, inner int
	objects := map[int]*strings.Builder{}
	idx := []int{0}

	for i, d := range data {
		if d == '{' {
			idx = append([]int{i}, idx[:]...)
			inner = 0
		}
		if d == '[' && idx[0] != 0 {
			inner++
		}
		if _, ok := objects[idx[0]]; !ok {
			objects[idx[0]] = &strings.Builder{}
		}
		if inner > 0 {
			objects[0].WriteRune(d)
		} else {
			objects[idx[0]].WriteRune(d)
		}
		if d == '}' {
			idx = idx[1:]
		}
		if d == ']' && idx[0] != 0 {
			inner--
		}
	}

	for k, v := range objects {
		str := v.String()
		if k != 0 && strings.Contains(str, "red") {
			continue
		}
		total += sumInts(str)
	}
	return total
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
