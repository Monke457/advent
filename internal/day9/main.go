package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := reader.FileToArray("data/day9.txt")

	values := [][]int{}
	for _, l := range lines {
		values = append(values, parseValues(l))
	}

	fmt.Println(solveFirstProblem(values))
	fmt.Println(solveSecondProblem(values))
}

func solveFirstProblem(values [][]int) int {
	result := 0
	for _, v := range values {
		result += getNext(v)
	}
	return result
}

func solveSecondProblem(values [][]int) int {
	result := 0
	for _, v := range values {
		result += getPrevious(v)
	}
	return result
}

func parseValues(line string) []int {
	vals := strings.Split(line, " ")

	result := []int{}
	for _, val := range vals {
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		result = append(result, v)
	}
	return result
}

func getNext(vals []int) int {
	diffs := []int{}
	done := true
	for i := 0; i < len(vals)-1; i++ {
		diff := vals[i+1] - vals[i]
		diffs = append(diffs, diff)
		if done && diff != 0 {
			done = false
		}
	}

	if done {
		return vals[len(vals)-1]
	}
	return getNext(diffs) + vals[len(vals)-1]
}

func getPrevious(vals []int) int {
	diffs := []int{}
	done := true
	for i := 0; i < len(vals)-1; i++ {
		diff := vals[i+1] - vals[i]
		diffs = append(diffs, diff)
		if done && diff != 0 {
			done = false
		}
	}

	if done {
		return vals[0]
	}
	return vals[0] - getPrevious(diffs)
}
